package main 

type PricingStrategy interface {
	calculatePrice(ticket *ParkingTicket) float64
}

type PerHourBasis struct {}

func (p *PerHourBasis) calculatePrice(ticket *ParkingTicket) float64 {
	duration := ticket.GetExitTime().Sub(ticket.GetEntryTime()).Milliseconds()

	switch ticket.vehicle.GetType() {
	case BIKE:
		return float64(duration) * 2 
	case CAR:
		return float64(duration) * 5 
	case VAN:
		return float64(duration) * 6 
	case TRUCK:
		return float64(duration) * 10 
	}
	return 0
}