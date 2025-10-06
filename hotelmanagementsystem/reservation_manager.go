package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ReservationManager struct {
	hotels       map[string]*Hotel
	rooms        map[string]*Room
	guests       map[string]*Guest
	reservations map[string]*Reservation
	// mapping between idempotency key and reservation ID 
	idempotency map[string]string

	mu sync.RWMutex 
}

var ReservationManagerInstance *ReservationManager
var ReservationManagerOnce sync.Once

func NewReservationManager() *ReservationManager {
	ReservationManagerOnce.Do(func() {
		ReservationManagerInstance = &ReservationManager{
			hotels: make(map[string]*Hotel),
			rooms: make(map[string]*Room),
			guests: make(map[string]*Guest),
			reservations: make(map[string]*Reservation),
			idempotency: make(map[string]string) ,
		}
	})

	return ReservationManagerInstance 
}


func (rm *ReservationManager)AddHotel(hotel *Hotel) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.hotels[hotel.id] = hotel 
}

func (rm *ReservationManager)AddRoom(room *Room) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.rooms[room.roomId] = room 
}

func (rm *ReservationManager)AddGuest(guest *Guest) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.guests[guest.id] = guest 
}

// Search returns available Rooms of roomType in a particular hotel for [start ,end)
func (rm *ReservationManager) Search(hotelID string, roomType RoomType, start time.Time, end time.Time) ([]string, error) {
	rm.mu.RLock()
	defer rm.mu.RUnlock() 

	roomIds := make([]string, 0)
	
	for _, room := range rm.rooms {
		if room.hotelId == hotelID && room.roomType == roomType && room.status == RoomStatus_Available && room.roomCalendar.IsAvailable(start, end) {
			roomIds = append(roomIds, room.roomId)
		}
	}
	return roomIds, nil

}

// SearchByCity finds available Rooms of RoomType in all hotels of a particular city for [start, end)
func (rm *ReservationManager) SearchByCity(city string, roomType RoomType, start time.Time, end time.Time) ([]string, error)  {
	rm.mu.RLock()
	defer rm.mu.RUnlock() 

	roomIds := make([]string, 0)

	for _, room := range rm.rooms {
		// match city 
		hotel := rm.hotels[room.hotelId]
		if hotel.city != city {
			continue; 
		}
		if room.roomType != roomType {
			continue 
		}
		if room.status != RoomStatus_Available {
			continue 
		}
		if room.roomCalendar.IsAvailable(start, end) {
			roomIds = append(roomIds, room.roomId)
		}
	} 

	return roomIds, nil 
}

// CreateHold finds a candidate room and places a HOLD. Idempotency supported via idempotencyKey.
func (rm *ReservationManager) CreateHold(hotelId string, guestId string, roomType RoomType, start time.Time, end time.Time, idempotencyKey string ) (*Reservation, error) {
	if !start.Before(end) {
		return nil, errors.New("invalid date range. ")
	}


	rm.mu.RLock()
	// check if idempotency key exists 
	reservationId, ok := rm.idempotency[idempotencyKey]
	if ok {
		// get the reservation corresponding to this idempotency key
		res, ok := rm.reservations[reservationId]
		if ok {
			return res, nil 
		}
	}
	rm.mu.RUnlock()

	// Get all available room Ids satisfying the criteria 
	candidateRoomIds, err := rm.Search(hotelId, roomType, start, end)
	if err != nil {
		return nil, err
	}

	if len(candidateRoomIds) == 0 {
		return nil, errors.New("no available rooms")
	}

	var chosenRoom *Room 
	reservationId = fmt.Sprintf("res:%s", uuid.NewString())

	// Try each candidate room id and attempt to Reserve(first-fit)
	for _, rID := range candidateRoomIds {
		rm.mu.RLock()
		room, ok := rm.rooms[rID]
		rm.mu.RUnlock()
		if !ok {
			continue; 
		}
		if room.status == RoomStatus_Available {
			err := room.roomCalendar.Reserve(reservationId, start, end)
			if err != nil {
				continue 
			}
			chosenRoom = room 
			break 
		}
	}

	if chosenRoom == nil {
		return nil, errors.New("no available rooms")
	}

	now := time.Now()

	// Build reservation object 
	reservation := &Reservation{
		reservationId: reservationId,
		hotelId: chosenRoom.hotelId,
		roomId: chosenRoom.roomId,
		roomType: chosenRoom.roomType,
		guestId: guestId,
		startDate: start,
		endDate: end,
		status: ReservationStatus_OnHold, 
		idempotency: idempotencyKey,
		totalAmount: computeTotalPrice(chosenRoom.basePrice, start, end),
		createdAt: now,
		updatedAt: now,
	}

	// store reservation and idempotency mapping
	rm.mu.Lock()
	rm.idempotency[idempotencyKey] = reservationId
	rm.reservations[reservationId] = reservation 
	rm.mu.Unlock() 

	return reservation, nil 
}

// CancelReservation marks a reservation cancelled 
func(rm *ReservationManager) CancelReservation(reservationId string) error {
	rm.mu.RLock()
	res, ok := rm.reservations[reservationId]
	rm.mu.RUnlock()
	if !ok {
		return fmt.Errorf("reservation not found")
	}
	// Already cancelled
	if res.status == ReservationStatus_Cancelled {
		return nil 
	}

	rm.mu.RLock()
	room := rm.rooms[res.roomId]
	rm.mu.RUnlock()
	
	// release calendar
	_ = room.roomCalendar.Release(reservationId)

	// set status to cancel 
	rm.mu.Lock()
	res.status = ReservationStatus_Cancelled
	res.updatedAt = time.Now() 
	rm.mu.Unlock()
	return nil 
}

// ConfirmReservation marks a reservation confirmed 
func(rm *ReservationManager) ConfirmReservation(reservationId string) error {
	rm.mu.RLock()
	res, ok := rm.reservations[reservationId]
	rm.mu.RUnlock()
	if !ok {
		return errors.New("reservation does not exist")
	}

	rm.mu.RLock()
	room := rm.rooms[res.roomId]
	rm.mu.RUnlock()

	// set room calendar status to confirmed
	_ = room.roomCalendar.Confirm(reservationId)

	// set reservation status to confirmed
	rm.mu.Lock()
	res.status = ReservationStatus_Confirmed
	res.updatedAt = time.Now() 
	rm.mu.Unlock()
	return nil 
}

// ProcessPayment processes payment using a payment strategy
func(rm *ReservationManager) ProcessPayment(reservationId string, paymentStrategy PaymentStrategy) error {
	rm.mu.RLock()
	res, ok := rm.reservations[reservationId]
	rm.mu.RUnlock()
	if !ok {
		return errors.New("reservation does not exist")
	}
	err := paymentStrategy.Pay(res.totalAmount) 
	if err != nil {
		// cancel reservation 
		_ = rm.CancelReservation(reservationId)
		return fmt.Errorf("payment failed. cancelling reservation")
	}

	// confirm booking 
	paymentObject := &Payment{
		id: fmt.Sprintf("pay:%s", uuid.NewString()),
		amount: res.totalAmount,
		createdAt: time.Now(),
	}
	res.payment = paymentObject 
	rm.ConfirmReservation(reservationId)

	return nil 
}

// Check in
func(rm *ReservationManager)CheckIn(guestId string, reservationId string) error {
	rm.mu.RLock()
	reservation, ok := rm.reservations[reservationId]
	rm.mu.RUnlock()
	if !ok {
		return fmt.Errorf("reservation %s does not exist", reservationId)
	}
	reservation.UpdateStatus(ReservationStatus_CheckedIn)
	return nil
}