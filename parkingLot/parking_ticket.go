package main 

import "time"


type ParkingTicket struct {
	id int 
	vehicle Vehicle 
	spot *ParkingSpot
	entryTime time.Time
	exitTime time.Time 
}

func NewParkingTicket(id int, vehicle Vehicle, spot *ParkingSpot, entryTime time.Time) *ParkingTicket {
	return &ParkingTicket{
		id: id, 
		vehicle: vehicle,
		spot: spot,
		entryTime: entryTime,
	}
}
func (pt *ParkingTicket) SetExitTime(exitTime time.Time ) {
	pt.exitTime = exitTime
}

func (pt *ParkingTicket) GetExitTime() time.Time {
	return pt.exitTime 
}


func (pt *ParkingTicket) GetEntryTime() time.Time {
	return pt.entryTime
}

func (pt *ParkingTicket) GetVehicle() Vehicle {
	return pt.vehicle
}


func (pt *ParkingTicket) GetSpot() *ParkingSpot {
	return pt.spot
}
