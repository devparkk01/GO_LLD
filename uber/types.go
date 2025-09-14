package main

type RideStatus string

type RideType string

const (
	RideStatusRequested   RideStatus = "Requested"
	RideStatusInProgress RideStatus = "InProgress"
	RideStatusCompleted RideStatus = "Completed"
	RideStatusCancelled   RideStatus = "Cancelled"
)


const (
	RideTypeRegular RideType = "Regular"
	RideTypePremium RideType = "Premium"
	RideTypeLuxury RideType = "Luxury"
)