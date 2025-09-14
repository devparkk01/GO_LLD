package main

import "sync"

type RiderManagementService struct {
	riders map[string]*Rider
	mu     sync.RWMutex
}

var riderServiceInstance *RiderManagementService
var riderServiceOnce sync.Once


func NewRiderManagementService() *RiderManagementService {
	riderServiceOnce.Do(func() {
		riderServiceInstance = &RiderManagementService{
			riders: make(map[string]*Rider),
		}
	})

	return riderServiceInstance
}


func(rs *RiderManagementService) AddRider(name string, contact int, location *Location) *Rider {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	rider := NewRider(name, contact, location)
	rs.riders[rider.id] = rider 
	return rider 
}


func(rs *RiderManagementService) GetRider(riderId string) *Rider {
	rs.mu.RLock()

	rider, ok := rs.riders[riderId]
	rs.mu.RUnlock()
	if !ok {
		return nil 
	}
	return rider 
}

func (rs *RiderManagementService) AddRideHistory(rider *Rider, ride *Ride) {
	rs.mu.Lock()
	defer rs.mu.Unlock() 

	rider.AddRideHistory(ride)
} 