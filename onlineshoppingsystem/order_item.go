package main

import "time"

type OrderItem struct {
	id           string
	orderId      string
	productId    string
	quantity     int
	pricePerUnit float32
	createdAt    time.Time
}