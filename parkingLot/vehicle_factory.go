package main 

type VehicleFactory struct {}

func (f *VehicleFactory) CreateVehicle(vType VehicleType, numberPlate string) Vehicle {
	return &BaseVehicle{
		vType: vType,
		numberPlate: numberPlate,
	}
}