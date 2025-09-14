package main 

type LimiterManager interface {
	// Returns all limiters as part of the limiter manager
	GetLimiters(req *RequestContext) []Limiter	
}


