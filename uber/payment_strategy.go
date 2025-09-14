package main

import (
	"fmt"
	"sync"
)

type PaymentService struct {}

var paymentServiceInstance *PaymentService 
var paymentServiceOnce sync.Once 

func NewPaymentService() *PaymentService {
	paymentServiceOnce.Do(func() {
		paymentServiceInstance = &PaymentService{}
	})
	return paymentServiceInstance
}

func(p *PaymentService)	ProcessPayment(amount float64, from string, to string, paymentStrategy PaymentStrategy) {
	paymentStrategy.ProcessPayment(amount, from, to)
}	


// --- Payment strategies 
type PaymentStrategy interface {
	ProcessPayment(amount float64, from string, to string)
}	

// implements PaymentStrategy
type CardPayment struct{}

func (c *CardPayment) ProcessPayment(amount float64, from string, to string) {
	fmt.Printf("Transferred amount %.2f from %s to %s via card.\n", amount, from, to)
}

// implements PaymentStrategy
type UPIPayment struct{}

func (c *UPIPayment) ProcessPayment(amount float64, from string, to string) {
	fmt.Printf("Transferred amount %.2f from %s to %s via UPI.\n", amount, from, to)
}