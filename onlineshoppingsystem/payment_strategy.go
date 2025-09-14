package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float32) error
}

type CardPayment struct{}

func (c *CardPayment) Pay(amount float32) error {
	fmt.Printf("Successfully paid %.2f by Card\n", amount)
	return nil 
}

type UPIPayment struct{}

func (c *UPIPayment) Pay(amount float32) error {
	fmt.Printf("Successfully paid %.2f by UPI\n", amount)
	return nil 
}