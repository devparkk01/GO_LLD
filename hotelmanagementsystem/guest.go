package main 


type Guest struct {
	id string 
	name string 
	email string 
}

func NewGuest(id, name, email string) *Guest {
	return &Guest {
		id: id,
		name: name, 
		email: email,
	}
}