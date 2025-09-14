package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) error
}

type CardPayment struct{}

func (c *CardPayment) Pay(amount float64) error {
	fmt.Printf("Successfully paid %.2f by Card\n", amount)
	return nil 
}

type UPIPayment struct{}

func (c *UPIPayment) Pay(amount float64) error {
	fmt.Printf("Successfully paid %.2f by UPI\n", amount)
	return nil 
}