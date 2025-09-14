package main 

type Vehicle struct {
	model string 
	noPlate string 
	rType RideType
}

func NewVehicle(model, noPlate string, rType RideType) *Vehicle {
	return &Vehicle{
		model: model,
		noPlate: noPlate,
		rType: rType,
	}
}