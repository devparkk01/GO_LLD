package main

import "github.com/google/uuid"

type Rider struct {
	id string 
	name string 
	contact int 
	location *Location 
	rideHistory []*Ride 
}

func NewRider(name string, contact int, location *Location) *Rider {
	return &Rider{
		id: uuid.NewString(),
		name: name, 
		contact: contact,
		location: location, 
	}
}

func(r *Rider) AddRideHistory(ride *Ride) {
	r.rideHistory = append(r.rideHistory, ride)
}