package main

import "fmt"

type User struct {
	id           string
	name         string
	email        string
	balanceSheet *BalanceSheet
}

func NewUser(id string, name string, email string) *User {
	return &User{
		id:           id,
		name:         name,
		email:        email,
		balanceSheet: NewBalanceSheet(),
	}
}

// returns total amount current user is Owed.
func (u *User) TotalIsOwed() float64 {
	return u.balanceSheet.TotalIsOwed()
}

// return total amount current user owes to others.
func (u *User) TotalOwed() float64 {
	return u.balanceSheet.TotalOwed()
}

// Get balance for OtherUserID 
func (u *User) GetBalance(otherUserId string) float64 {
	return u.balanceSheet.GetBalance(otherUserId)
}

// Settle balance for OtherUserID 
func (u *User) SettleBalance(otherUserID string) {
	amount := u.GetBalance(otherUserID)
	u.balanceSheet.SettleBalance(otherUserID)
	fmt.Printf("Settled amount %.2f for %s with %s \n", amount, u.id, otherUserID)
}

// Updates balance sheet of the current user, for otherUserid
func (u *User) UpdateBalanceSheet(otherUserid string, amount float64) {
	u.balanceSheet.UpdateBalance(otherUserid, amount)
}
