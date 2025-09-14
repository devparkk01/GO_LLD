package main

import (
	"fmt"
	"time"
)


type ExitGate struct {
	id int 
}

func NewExitGate(id int) *ExitGate {
	return &ExitGate{id: id }
}

func (eg *ExitGate) ProcessTicket(ticket *ParkingTicket, pricing PricingStrategy ,payment PaymentStrategy) bool {
	// set exit time 
	ticket.SetExitTime(time.Now())
	// calculate price
	amount := pricing.calculatePrice(ticket)
	err := payment.Pay(amount)
	if err != nil {
		fmt.Printf("Error while making the payment, %s\n", err.Error()) 
		return false 
	}
	
	fmt.Printf("Exit gate: %d: vehicle %s exited.\n", eg.id, ticket.GetVehicle().GetNumberPlate() )
	ticket.GetSpot().Unpark()
	return true 
	
}