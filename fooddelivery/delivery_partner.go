package main 


type DeliveryPartner struct {
	id string 
	name string 
	phone string 
	isAvailable bool // true means driver is available to fulfill an order, false means either driver is unavailable or he's
	                  // fulfilling an existing order 
	address *Address
}

func NewDeliveryPartner(id, name, phone string, isAvailable bool, address *Address) *DeliveryPartner {
	return &DeliveryPartner{
		id: id, 
		name: name, 
		phone: phone, 
		isAvailable: isAvailable,
		address: address,
	}
}

func(d *DeliveryPartner) SetAvailability(availability bool) {
	d.isAvailable = availability
}

func(d *DeliveryPartner) checkAvailable() bool {
	return d.isAvailable 
}