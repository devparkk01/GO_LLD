package main

import (
	"fmt"
	"sync"
)

var parkingLotInstance *ParkingLot 

type ParkingLot struct {
	floors []*ParkingFloor
	mu sync.Mutex // to handle concurrent vehicles trying to find a parking spot 
}

func GetParkingLotInstance() *ParkingLot {
	if parkingLotInstance == nil {
		parkingLotInstance = &ParkingLot{}
	}
	return parkingLotInstance
}

func (pl *ParkingLot) AddFloor(floor *ParkingFloor) {
	pl.floors = append(pl.floors, floor)
}

func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) *ParkingSpot{
	pl.mu.Lock()
	defer pl.mu.Unlock()

	for _, floor := range pl.floors {
		for _, spot := range floor.parkingSpots {
			if spot.CanPark(vehicle) {
				// park this vehicle here 
				spot.ParkVehicle(vehicle)
				fmt.Printf("Vehicle %s parked successfully in spot %d on floor %d\n", vehicle.GetNumberPlate(), spot.GetId(), floor.GetFloorNo())
				return spot 
			}
		}
	}
	fmt.Printf("Unable to find parking spot for vehicle %s\n", vehicle.GetNumberPlate())
	return nil 
}

func (pl *ParkingLot) VacateSpot(spot *ParkingSpot) {
	if spot.IsFree() {
		fmt.Printf("Spot is already vacated\n")
		return 
	}
	spot.Unpark()
	fmt.Printf("Spot %d has been vacated. \n", spot.GetId() )
}

func (pl *ParkingLot) DisplayAvailability() {
	for _, floor := range pl.floors {
		floor.DisplayAvailability()
	}
}