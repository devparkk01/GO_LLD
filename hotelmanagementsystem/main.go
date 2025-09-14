package main

import (
	"fmt"
	"time"
)

func main() {
	rm := NewReservationManager()
	h1 := NewHotel("hotel-1", "Ocean view", "Delhi")
	h2 := NewHotel("hotel-2", "city center", "Mumbai")
	h3 := NewHotel("hotel-3", "Big Suite", "Delhi")

	rm.AddHotel(h1)
	rm.AddHotel(h2)
	rm.AddHotel(h3)

	room1 := NewRoom("room-101", "101", h1.id, ROOMTYPE_DOUBLE, RoomStatus_Available, 10000, 3)
	room2 := NewRoom("room-102", "102", h1.id, ROOMTYPE_DOUBLE, RoomStatus_Available, 8000, 3)
	room3 := NewRoom("room-103", "103", h1.id, ROOMTYPE_FAMILY, RoomStatus_Available, 12000, 4)
	room4 := NewRoom("room-104", "104", h1.id, ROOMTYPE_FAMILY, RoomStatus_OutOfService, 12500, 4)
	room5 := NewRoom("room-201", "201", h2.id, ROOMTYPE_LUXURY, RoomStatus_Available, 25000, 6)
	room6 := NewRoom("room-202", "202", h2.id, ROOMTYPE_FAMILY, RoomStatus_Available, 18000, 4)
	room7 := NewRoom("room-501", "501", h3.id, ROOMTYPE_DOUBLE, RoomStatus_Available, 7500, 3)

	rm.AddRoom(room1)
	rm.AddRoom(room2)
	rm.AddRoom(room3)
	rm.AddRoom(room4)
	rm.AddRoom(room5)
	rm.AddRoom(room6)
	rm.AddRoom(room7)

	g1 := NewGuest("guest-1", "Alice", "alice@gmail.com")
	g2 := NewGuest("guest-2", "Bob", "bob@gmail.com")
	rm.AddGuest(g1)
	rm.AddGuest(g2)

	start := normalizeDate(time.Now().AddDate(2025, 9, 10))
	end := normalizeDate(time.Now().AddDate(2025, 9, 13))
	
	roomIDs, _ := rm.SearchByCity("Delhi", ROOMTYPE_DOUBLE, start, end )
	fmt.Println(roomIDs)

	roomIDs, _ = rm.Search(h1.id, ROOMTYPE_DOUBLE, start, end)
	fmt.Println(roomIDs)


	res, err := rm.CreateHold(h1.id, g1.id, ROOMTYPE_DOUBLE, start, end, "idem:12345")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	err = rm.ProcessPayment(res.reservationId, &CardPayment{})
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(res)

	roomIDs, _ = rm.Search(h1.id, ROOMTYPE_DOUBLE, start, end)
	fmt.Println(roomIDs)


	res, err = rm.CreateHold(h1.id, g2.id, ROOMTYPE_DOUBLE, start, end, "idem:34566")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	err = rm.ProcessPayment(res.reservationId, &CardPayment{})
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(res)

	roomIDs, _ = rm.Search(h1.id, ROOMTYPE_DOUBLE, start, end)
	fmt.Println(roomIDs)


}
