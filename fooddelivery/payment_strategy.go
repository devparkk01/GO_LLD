package main

import "fmt"

type PaymentStrategy interface {
	ProcessPayment(amount float32) error
}

type UPIPayment struct{}

func (u *UPIPayment) ProcessPayment(amount float32) error {
	fmt.Printf("successfully proccessed payment of amount %.2f via UPI.\n", amount)
	return nil
}

type CardPayment struct{}

func (c *CardPayment) ProcessPayment(amount float32) error {
	fmt.Printf("successfully proccessed payment of amount %.2f via card.\n", amount)
	return nil
}