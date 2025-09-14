package main 

type SpotType int 

const (
	SMALL SpotType = iota
	COMPACT 
	LARGE 
)

type ParkingSpot struct{
	id int 
	vehicle Vehicle
	sType SpotType 
}

func NewParkingSpot(id int, sType SpotType) *ParkingSpot {
	return &ParkingSpot{
		id: id, 
		vehicle: nil,
		sType: sType,
	}
}

func(p *ParkingSpot) GetId() int {
	return p.id
}

func(p *ParkingSpot) GetVehicle() Vehicle {
	return p.vehicle 
}

func(p *ParkingSpot) GetSpotType() SpotType {
	return p.sType
}

func (p *ParkingSpot) IsFree() bool {
	return p.vehicle == nil 
}

func (p *ParkingSpot) CanPark(v Vehicle) bool {
	vehicleType := v.GetType()
	switch p.GetSpotType(){
	case SMALL: 
		return vehicleType == BIKE || vehicleType == CAR 
	case COMPACT:
		return vehicleType == VAN
	case LARGE:
		return vehicleType == TRUCK
	default:
		return false 
	}
}

func (p *ParkingSpot) ParkVehicle(vehicle Vehicle) bool {
	if p.IsFree() && p.CanPark(vehicle) {
		p.vehicle = vehicle
		return true 
	}
	return false  
}

func (p *ParkingSpot) Unpark() {
	p.vehicle = nil 
}

type ParkingSpotFactory struct{}

func (ps *ParkingSpotFactory) CreateSpot(id int, sType SpotType) *ParkingSpot {
	return &ParkingSpot{
		id: id, 
		sType: sType,
	}
}