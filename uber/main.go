package main

import (
	"time"
)


func main() {
	driverService := NewDriverManagementService()
	riderService := NewRiderManagementService() 
	paymentService := NewPaymentService()
	rideService := NewRideManagementService(driverService, riderService, paymentService)


	// Add drivers to the system 
	driver1 := driverService.AddDriver("Phil", NewVehicle("Toyota", "KA3456", RideTypePremium), NewLocation(25, 30))
	driver2 := driverService.AddDriver("Tony", NewVehicle("Maruti", "KA0923", RideTypeRegular), NewLocation(47, 19))
	driver3 := driverService.AddDriver("Ezequiel", NewVehicle("BYD", "KA1094", RideTypePremium), NewLocation(10, 38))
	driver4 := driverService.AddDriver("Roy", NewVehicle("TATA", "KA743232", RideTypeRegular), NewLocation(10, 38))
	driver5 := driverService.AddDriver("Peter", NewVehicle("BYD", "KA03298", RideTypeRegular), NewLocation(66, 24))

	driverService.SetAvailability(driver1.GetId(), true)
	driverService.SetAvailability(driver2.GetId(), true)
	driverService.SetAvailability(driver3.GetId(), true)
	driverService.SetAvailability(driver4.GetId(), true)
	driverService.SetAvailability(driver5.GetId(), true)


	// Add riders to the system 
	rider1 := riderService.AddRider("Harry", 123452312, NewLocation(24, 36))
	rider2 := riderService.AddRider("Shane", 8976893, NewLocation(88, 45))

	ride := rideService.RequestRide(rider1.id, NewLocation(11, 45), NewLocation(45, 123), RideTypeRegular)
	if ride == nil {
		return 
	}
	secondRide := rideService.RequestRide(rider2.id, NewLocation(25, 87), NewLocation(1, 43), RideTypeRegular)
	if secondRide == nil {
		return 
	}


	// wait to simulate driver coming at pickup 
	time.Sleep(time.Second * 1)
	rideService.StartRide(ride)
	rideService.StartRide(secondRide)

	time.Sleep(time.Second * 3)
	rideService.EndRide(ride, &CardPayment{})
	time.Sleep(time.Second * 1)
	rideService.EndRide(secondRide, &UPIPayment{})

}