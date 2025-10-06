package main 

type OrderStatus int 

const (
	OrderStatusPending OrderStatus = iota
	OrderStatusCancelled
	OrderStatusConfirmed
	OrderStatusPreparing 
	OrderStatusOutForDelivery
	OrderStatusDelivered
)
