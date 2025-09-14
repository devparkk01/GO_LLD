package main

import "time"

type Ride struct {
	id         string
	rider      *Rider
	driver     *Driver
	rideStatus RideStatus
	fare       float64
	pickup     *Location
	drop       *Location
	rideType   RideType
	startTime  time.Time
	endTime    time.Time
}
