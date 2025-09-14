package main

import "time"

type PaymentStatus int 

const (
	PaymentStatus_Success PaymentStatus = iota
	PaymentStatus_Failure 
	PaymentStatus_Pending
)

type Order struct {
	id            string
	userId        string
	totalPrice    float32
	paymentStatus PaymentStatus
	items         []*OrderItem
	createdAt     time.Time
}