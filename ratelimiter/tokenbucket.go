package main

import (
	// "fmt"
	"fmt"
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	name           string     // name of the token bucket
	maxCapacity    int        // max capacity this bucket can hold
	tokens         int        // tokens available in the bucket at present
	refillRate     int        // refill rate of this token bucket  tokens/sec
	lastRefillTime time.Time  // last refill time
	mu             sync.Mutex // mutex for concurrent access
}

func NewTokenBucket(name string, maxCapacity int, refillRate int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		name:           name,
		maxCapacity:    maxCapacity,
		tokens:         maxCapacity,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (t *TokenBucketLimiter) Allow(req *RequestContext) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// // Following is the logic for refilling
	// now := time.Now()
	// // check when was the last refill
	// if now.Sub(t.lastRefillTime) >= time.Second {
	// 	// refill the bucket
	// 	t.tokens = t.tokens
	// 	t.lastRefillTime = now
	// }

	t.Refill()

	// if we have tokens left in the bucket
	if t.tokens > 0 {
		t.tokens--
		return true
	}
	return false
}

func (t *TokenBucketLimiter) Refill() {
	now := time.Now()
	// calculate the difference between last refill time
	elapsedTimeInSecond := now.Sub(t.lastRefillTime).Seconds()
	// count new tokens to be added
	newTokens := int(elapsedTimeInSecond) * t.refillRate
	if newTokens > 0 {
		// Add the new tokens
		fmt.Println("refilling")
		t.tokens += newTokens
		t.lastRefillTime = now
	}
	if t.tokens > t.maxCapacity {
		t.tokens = t.maxCapacity
	}
}

func (t *TokenBucketLimiter) Name() string {
	return t.name
}
