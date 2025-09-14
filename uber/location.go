package main

import "math"

type Location struct {
	x, y int
}

func NewLocation(x, y int) *Location {
	return &Location{
		x: x,
		y: y,
	}
}

func (l *Location) ToDistance(location *Location) float64 {
	return math.Sqrt(float64((l.x - location.x) * (l.x - location.x) + (l.y - location.y) * (l.y - location.y)))
}