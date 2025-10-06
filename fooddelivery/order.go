package main

import "time"

type Order struct {
	id              string
	customerId      string
	restaurantId    string
	items           []*OrderItem
	orderStatus     OrderStatus
	deliveryPartner *DeliveryPartner
	amount         float32
	createdAt       time.Time
}

