package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type RideManagementService struct {
	rides          map[string]*Ride
	driverService  *DriverManagementService
	riderService   *RiderManagementService
	paymentService *PaymentService
	rideBaseFare   map[RideType]float64
	mu sync.RWMutex
}

var rideServiceInstance *RideManagementService
var rideServiceOnce sync.Once 

func NewRideManagementService(driverService *DriverManagementService, riderService *RiderManagementService, paymentService *PaymentService) *RideManagementService {
	rideServiceOnce.Do(func() {
		rideServiceInstance = &RideManagementService{
			rides: make(map[string]*Ride),
			driverService: driverService,
			riderService: riderService,
			paymentService: paymentService,
			rideBaseFare: map[RideType]float64{
				RideTypeRegular: 25,
				RideTypePremium: 50,
				RideTypeLuxury: 100,
			},
		}
	})

	return rideServiceInstance
}

func (r *RideManagementService) RequestRide(riderId string, pickup *Location, drop *Location, rType RideType) *Ride{
	/* Logic: 
	1. Get nearest available driver ( call driverService )
	2. Create a ride 
	3. Assign this driver to the ride
	4. Set availability of the driver to false 
	5. Add the ride to both rider's and driver's  history 

	*/
	r.mu.Lock()
	defer r.mu.Unlock()

	rider := r.riderService.GetRider(riderId)
	if rider == nil {
		fmt.Println("Rider not found ", riderId)
		return nil 
	}

	driver := r.driverService.FindNearbyDriver(pickup, rType)
	if driver == nil {
		fmt.Println("No available drivers")
		return nil 
	}

	// create a ride
	ride := &Ride{
		id: fmt.Sprintf("ride-%s", uuid.NewString()),
		rider: rider,
		driver: driver,
		pickup: pickup,
		drop: drop, 
		rideType: rType,
		rideStatus: RideStatusRequested,
	}
	r.rides[ride.id] = ride 

	fmt.Printf("Ride accepted for %s. Driver %s is coming at pickup location shortly\n" , rider.name, driver.GetName())

	// Driver accepts the ride 
	r.driverService.AcceptRide(driver, ride)

	// Add the ride to the ride history 
	r.riderService.AddRideHistory(rider, ride)

	return ride 
}


func( r *RideManagementService) StartRide(ride *Ride) {
	ride.startTime = time.Now()
	ride.rideStatus = RideStatusInProgress
	fmt.Printf("Ride for %s has started\n", ride.rider.name)
}

func( r *RideManagementService) EndRide(ride *Ride, paymentStrategy PaymentStrategy) {
	// calculate endtime 
	ride.endTime = time.Now()
	// calculate fare
	fare := r.CalculateFare(ride)
	ride.fare = fare 

	fmt.Printf("Total fare amount %.2f\n", fare)
	// Process payment 
	r.paymentService.ProcessPayment(fare, ride.rider.name, ride.driver.name, paymentStrategy)

	// Set driver's availability to true 

	r.driverService.SetAvailability(ride.driver.id, false)
}

func(r *RideManagementService) CalculateFare(ride *Ride) float64 {
	// get the base fare for ride type 
	baseFare := r.rideBaseFare[ride.rideType]
	// calculate distance 
	dist := ride.pickup.ToDistance(ride.drop)
	// calculate time 
	duration := ride.endTime.Sub(ride.startTime).Seconds()

	fare := baseFare + dist * 10 + duration * 5
	return fare
}