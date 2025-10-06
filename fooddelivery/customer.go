package main 

type Customer struct {
	id string 
	name string 
	phone string 
	email string 
	address *Address
}


func NewCustomer(id, name, phone, email string, address *Address) *Customer {
	return &Customer{
		id: id,
		name: name, 
		phone: phone,
		email: email, 
		address: address,
	}
}