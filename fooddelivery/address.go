package main 

type Address struct {
	street string 
	city string 
	pinCode string 
	latitude float32
	longtitude float32 
}

func NewAddress(street, city, pinCode string, latitude, longtitude float32) *Address {
	return &Address{
		street: street, 
		city: city,
		pinCode: pinCode,
		latitude: latitude,
		longtitude: longtitude,
	}
}