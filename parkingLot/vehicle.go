package main 

type VehicleType int 

const (
	BIKE VehicleType = iota 
	CAR 
	VAN 
	TRUCK 
)

type Vehicle interface {
	GetType() VehicleType
	GetNumberPlate() string  
}

type BaseVehicle struct {
	vType VehicleType
	numberPlate string 
}

func (v *BaseVehicle) GetType() VehicleType {
	return v.vType
}

func (v *BaseVehicle) GetNumberPlate() string {
	return v.numberPlate
}