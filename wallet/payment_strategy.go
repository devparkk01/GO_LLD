package main

import "fmt"

type PaymentStrategy interface {
	pay(amount float64, currency Currency) bool
}

type UPIPayment struct{}
type CardPayment struct{}
type BankTransfer struct{}

func (u *UPIPayment) pay(amount float64, currency Currency) bool {
	fmt.Printf("Processing UPI payment of %.2f %s...\n", amount, currency.code )
	return true 
}
func (c *CardPayment) pay(amount float64, currency Currency) bool {
	fmt.Printf("Processing Card payment of %.2f %s...\n", amount, currency.code )
	return true 
}
func (b *BankTransfer) pay(amount float64, currency Currency) bool {
	fmt.Printf("Processing Bank transfer payment of %.2f %s...\n", amount, currency.code )
	return true 
}
