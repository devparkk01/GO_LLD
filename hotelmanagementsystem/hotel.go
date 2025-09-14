package main


type Hotel struct {
	id string 
	name string 
	city string 
}

func NewHotel(id, name, city string) *Hotel {
	return &Hotel{
		id: id, 
		name: name,
		city: city,
	}
}