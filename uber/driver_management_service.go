package main

import (
	"fmt"
	"math"
	"sync"
)

type DriverManagementService struct {
	drivers map[string]*Driver
	mu      sync.RWMutex
}

var driverServiceInstance *DriverManagementService
var driverServiceOnce sync.Once 

func NewDriverManagementService() *DriverManagementService {
	driverServiceOnce.Do(func() {
		driverServiceInstance = &DriverManagementService{
			drivers: make(map[string]*Driver),
		}
	})
	return driverServiceInstance
}

func (ds *DriverManagementService) AddDriver(name string, vehicle *Vehicle, location *Location) *Driver{
	ds.mu.Lock() 
	defer ds.mu.Unlock()

	driver := NewDriver(name, vehicle, location)
	ds.drivers[driver.id] = driver
	return driver 
}

func (ds *DriverManagementService) FindNearbyDriver(pickup *Location, rType RideType) *Driver {
	/* loop through all drivers
		check if a driver is available, and has a vehicle of rType, calculate it's current location from pickup 
	*/
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	minDist := math.MaxFloat64
	var nearestDriver *Driver 


	for _, driver := range ds.drivers {
		if driver.GetAvailability() && driver.GetVehicleType() == rType {
			distance := driver.GetLocation().ToDistance(pickup)
			fmt.Println("Distance ", driver.name, ": ", distance)
			if distance < minDist {
				nearestDriver = driver 
				minDist = distance
			}
		}
	}
	return nearestDriver 
}

func(ds *DriverManagementService) SetAvailability(driverId string, availability bool) {
	ds.mu.RLock()
	driver, ok := ds.drivers[driverId]
	defer ds.mu.RUnlock()
	if !ok {
		return 
	}
	driver.SetAvailability(availability)
}

func(ds *DriverManagementService) AcceptRide(driver *Driver, ride *Ride) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	driver.AddRideHistory(ride)
	driver.SetAvailability(false)
}