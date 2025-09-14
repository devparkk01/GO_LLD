package main

import (
	"fmt"
	"time"
)

func main() {
	manager := NewTokenBucketManager(5, 10, 4, 2, 5, 2)

	rateLimiterService := RateLimiterEngine{
		limiterManager: manager,
	}

	requests := make([]*RequestContext, 50)

	for i := range requests {
		requests[i] = &RequestContext{id: i, userId: "user1", ip: "192.168.1.1", endpoint: "/api/data"}
	}

	fmt.Println(requests)

	func() {
		for _, req := range requests[:10] {
			go rateLimiterService.Allow(req)
		}
	}()

	time.Sleep(1 * time.Second) // give time for refilling

	func() {
		for _, req := range requests[10:] {
			go rateLimiterService.Allow(req)
		}
	}()

	time.Sleep(4 * time.Second)
}