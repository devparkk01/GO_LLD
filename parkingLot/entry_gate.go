package main 

import (
	"time"
)

type EntryGate struct {
	id int 
}

func NewEntryGate(id int ) *EntryGate {
	return &EntryGate{
		id: id, 
	}
}

func (eg *EntryGate) IssueTicket(vehicle Vehicle, parkingLot *ParkingLot) *ParkingTicket {
	spot := parkingLot.ParkVehicle(vehicle)
	if spot == nil {
		return nil
	}
	ticket := NewParkingTicket(12, vehicle, spot, time.Now())
	return ticket 
}