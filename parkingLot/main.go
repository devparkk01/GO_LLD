package main

import (
	"fmt"
	"time"
)

func main() {
	lot := GetParkingLotInstance()

	// setup floor and spots
	spotFactory := &ParkingSpotFactory{}
	spots := []*ParkingSpot{
		spotFactory.CreateSpot(11, COMPACT),
		spotFactory.CreateSpot(12, COMPACT),
		spotFactory.CreateSpot(13, LARGE),
		spotFactory.CreateSpot(14, LARGE),
	}
	firstFloor := NewParkingFloor(1, spots)

	secondFloor := NewParkingFloor(2, []*ParkingSpot{
		spotFactory.CreateSpot(21, COMPACT),
		spotFactory.CreateSpot(22, LARGE),
		spotFactory.CreateSpot(23, SMALL),
		spotFactory.CreateSpot(24, SMALL),
		spotFactory.CreateSpot(25, SMALL),
	})

	lot.AddFloor(firstFloor)
	lot.AddFloor(secondFloor)

	// setup entry and exit gate
	entryGate := NewEntryGate(1)
	exitGate := NewExitGate(1)

	lot.DisplayAvailability()

	// setup vehicle
	vehicleFactory := VehicleFactory{}
	firstCar := vehicleFactory.CreateVehicle(CAR, "KA-234234")
	// secondCar := vehicleFactory.CreateVehicle(CAR, "KA-23129")
	// firstBike := vehicleFactory.CreateVehicle(BIKE, "KA-001122")
	// secondBike := vehicleFactory.CreateVehicle(BIKE, "KA-112233")

	ticket := entryGate.IssueTicket(firstCar, lot)
	time.Sleep(time.Second * 2)
	if ticket != nil {
		creditCardPayment := &CreditCardPayment{}
		perHourPrice := &PerHourBasis{}
		if exitGate.ProcessTicket(ticket, perHourPrice, creditCardPayment) {
			fmt.Println("Exit succesful")
		}
	}
	

}