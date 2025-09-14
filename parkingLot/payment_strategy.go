package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) error 
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Printf("Amount %v has been successfully paid\n", amount)
	return nil 
}

type CashPayment struct{}

func (c *CashPayment) Pay(amount float64) error{
	fmt.Printf("Amount %v has been successfully paid\n", amount)
	return nil 
}


type UPIPayment struct{}

func (c *UPIPayment) Pay(amount float64) error {
	fmt.Printf("Amount %v has been successfully paid\n", amount)
	return nil 
}

