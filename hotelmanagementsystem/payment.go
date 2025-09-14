package main

import "time"

type Payment struct {
	id        string
	amount    float64
	createdAt time.Time
}