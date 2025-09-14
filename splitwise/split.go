package main 


type Split struct {
	user *User  
	amount float64
}


func NewSplit(user *User, amount float64) *Split {
	return &Split{user: user, amount: amount}
}