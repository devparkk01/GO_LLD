package main 

type MenuItem struct {
	id string 
	name string 
	price float32 
	isAvailable bool 
}

func NewMenuItem(id, name string, price float32, isAvailable bool) *MenuItem {
	return &MenuItem{
		id: id, 
		name: name, 
		price: price, 
		isAvailable: isAvailable,
	}
}