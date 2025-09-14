package main 


type User struct {
	id string 
	name string 
	wallet *Wallet 
}


func NewUser(id , name string) *User {
	return &User{
		id: id,
		name: name,
		wallet: NewWallet(id),
	}
}


func (u *User) GetWallet() *Wallet{
	return u.wallet
}

func (u *User) GetId() string{
	return u.id
}
