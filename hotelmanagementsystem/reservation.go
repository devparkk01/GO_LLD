package main

import "time"

type Reservation struct {
	reservationId        string
	hotelId   string
	roomId    string
	guestId   string
	roomType  RoomType
	startDate time.Time
	endDate time.Time 
	status ReservationStatus
	idempotency string 

	payment *Payment
	totalAmount float64
	createdAt time.Time 
	updatedAt time.Time 
}