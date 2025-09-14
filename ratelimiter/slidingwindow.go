package main

import (
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	name     string
	capacity int
	window   time.Duration
	requests []time.Time 
	mu sync.Mutex  
}

func NewSlidingWindowLimiter(name string, capacity int, window time.Duration ) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		name: name,
		capacity: capacity,
		window: window,
		requests: make([]time.Time, 0, capacity),
	}
}

func (l *SlidingWindowLimiter) Allow(req *RequestContext) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	cutoffTime := now.Add(-l.window) // get the cutoff time by subtracting window time from current time 

	// remove expired items from the front 
	for len(l.requests) > 0 && cutoffTime.After(l.requests[0]) {
		l.requests = l.requests[1:] // O(1) ammortized 
	}

	if len(l.requests) < l.capacity {
		l.requests = append(l.requests, now)
		return true 
	}
	return false 
	
}

func (l *SlidingWindowLimiter) Name() string {
	return l.name 
}