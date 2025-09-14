package main

import "fmt"

// contains limiterManager
type RateLimiterEngine struct {
	limiterManager LimiterManager
}

func(t *RateLimiterEngine) Allow(req *RequestContext) bool {
	for _, limiter := range t.limiterManager.GetLimiters(req) {
		if !limiter.Allow(req) {
			fmt.Println("Request ID: ", req.id , " has been rate limited.")
			return false 
		}
	}
	fmt.Println("Requst ID: ", req.id, " has been allowed.")
	return true 
}