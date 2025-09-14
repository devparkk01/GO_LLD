package main

import "github.com/google/uuid"


type Driver struct {
	id string 
	name string 
	vehicle *Vehicle 
	location *Location 
	isAvailable bool 
	rideHistory []*Ride 
}

func NewDriver(name string, vehicle *Vehicle, location *Location) *Driver {
	return &Driver{
		id: uuid.NewString(), 
		name: name,
		vehicle: vehicle,
		location: location,
		isAvailable: false,
	}
}

func (d *Driver) AddVehicle(vehicle *Vehicle) {
	d.vehicle = vehicle
}

func (d *Driver) UpdateLocation(location *Location) {
	d.location = location
}

func (d *Driver) AddRideHistory(ride *Ride) {
	d.rideHistory = append(d.rideHistory, ride)
}

func (d *Driver) SetAvailability(availability bool) {
	d.isAvailable = availability
}


// Getter methods 
func (d *Driver) GetId() string {
	return d.id
}

func (d *Driver) GetName() string {
	return d.name
}

func (d *Driver) GetAvailability() bool {
	return d.isAvailable
}

func (d *Driver) GetVehicle() *Vehicle {
	return d.vehicle
}

func (d *Driver) GetLocation() *Location {
	return d.location
}

func (d *Driver) GetVehicleType() RideType {
	return d.vehicle.rType
}