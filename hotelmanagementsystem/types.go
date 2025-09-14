package main

type RoomType string

const (
	ROOMTYPE_SINGLE RoomType = "Single"
	ROOMTYPE_DOUBLE RoomType = "Double"
	ROOMTYPE_FAMILY RoomType = "Family"
	ROOMTYPE_LUXURY RoomType = "Luxury"
)

type RoomStatus string

const (
	RoomStatus_Available    RoomStatus = "AVAILABLE"
	RoomStatus_Occupied     RoomStatus = "Occupied"
	RoomStatus_Dirty        RoomStatus = "Dirty"
	RoomStatus_OutOfService RoomStatus = "Out of Service"
)

type ReservationStatus string

const (
	ReservationStatus_Confirmed  ReservationStatus = "Confirmed"
	ReservationStatus_OnHold     ReservationStatus = "OnHold"
	ReservationStatus_CheckedIn  ReservationStatus = "CheckedIn"
	ReservationStatus_CheckedOut ReservationStatus = "CheckedOut"
	ReservationStatus_Cancelled  ReservationStatus = "Cancelled"
)
