package main

import "fmt"

type ParkingFloor struct {
	floorNo      int
	parkingSpots []*ParkingSpot
}

func NewParkingFloor(floorNo int, spots []*ParkingSpot) *ParkingFloor {
	return &ParkingFloor{
		floorNo:      floorNo,
		parkingSpots: spots,
	}
}

func (pf *ParkingFloor) GetParkingSpots() []*ParkingSpot {
	return pf.parkingSpots
}

func (pf *ParkingFloor) GetFloorNo() int {
	return pf.floorNo
}

func (pf *ParkingFloor) ParkVehicle(v Vehicle) bool {
	for _, spot := range pf.GetParkingSpots() {
		if spot.IsFree() && spot.CanPark(v) {
			return spot.ParkVehicle(v)
		}
	}
	return false
}

func (pf *ParkingFloor) UnParkVehicle(v Vehicle) bool {
	for _, spot := range pf.GetParkingSpots() {
		if spot.GetVehicle() == v {
			spot.Unpark()
			return true
		}
	}
	return false
}

func (pf *ParkingFloor) DisplayAvailability() {
	for _, spot := range pf.GetParkingSpots() {
		status := "Available"
		if !spot.IsFree() {
			status = "Occupied"
		}
		fmt.Println("Level: ", pf.GetFloorNo(), " Spot: ", spot.GetId(), " Spot Type: ", spot.GetSpotType(), " Status: ", status)
	}
}