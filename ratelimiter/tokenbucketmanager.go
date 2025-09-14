package main

import (
	// "fmt"
	"sync"
)

type TokenBucketManager struct {
	userBuckets     map[string]*TokenBucketLimiter
	ipBuckets       map[string]*TokenBucketLimiter
	endpointBuckets map[string]*TokenBucketLimiter

	userCapacity     int
	ipCapacity       int
	endpointCapacity int

	userRefillRate     int
	ipRefillRate       int
	endpointRefillRate int

	mu sync.Mutex
}

func NewTokenBucketManager(userCap, ipCap, endpointCap, userRefill, ipRefill, endpointRefill int) *TokenBucketManager {
	return &TokenBucketManager{
		userBuckets:     make(map[string]*TokenBucketLimiter),
		ipBuckets:       make(map[string]*TokenBucketLimiter),
		endpointBuckets: make(map[string]*TokenBucketLimiter),

		userCapacity:     userCap,
		ipCapacity:       ipCap,
		endpointCapacity: endpointCap,

		userRefillRate:     userRefill,
		ipRefillRate:       ipRefill,
		endpointRefillRate: endpointRefill,
	}
}

// Gets IpTokenBucket for the given endpoint if it exists, else Create a new IpTokenBucket
func (t *TokenBucketManager) GetOrCreateIPTokenBucket(ip string) Limiter {
	t.mu.Lock()
	defer t.mu.Unlock()
	ipBucket, exists := t.ipBuckets[ip]
	if !exists {
		ipBucket = NewTokenBucket("ip_token_bucket"+ip, t.ipCapacity, t.ipRefillRate)
		t.ipBuckets[ip] = ipBucket
	}
	return ipBucket
}

// Gets UserTokenBucket for the given endpoint if it exists, else Create a new UserTokenBucket
func (t *TokenBucketManager) GetOrCreateUserTokenBucket(user string) Limiter {
	t.mu.Lock()
	defer t.mu.Unlock()
	userBucket, exists := t.userBuckets[user]
	if !exists {
		userBucket = NewTokenBucket("user_token_bucket"+user, t.userCapacity, t.userRefillRate)
		t.userBuckets[user] = userBucket
	}
	return userBucket
}

// Gets EndpointTokenBucket for the given endpoint if it exists, else Create a new EndpointTokenBucket
func (t *TokenBucketManager) GetOrCreateEndpointTokenBucket(endpoint string) Limiter {
	t.mu.Lock()
	defer t.mu.Unlock()
	endpointBucket, exists := t.endpointBuckets[endpoint]
	if !exists {
		endpointBucket = NewTokenBucket("endpoint_token_bucket"+endpoint, t.endpointCapacity, t.endpointRefillRate)
		t.endpointBuckets[endpoint] = endpointBucket
	}
	return endpointBucket
}

// returns all limiters as part of this Limiter manager
func (t *TokenBucketManager) GetLimiters(req *RequestContext) []Limiter {
	// get the userBucket for this key (userID )
	userTokenBucket := t.GetOrCreateUserTokenBucket(req.userId)
	// get the ipBucket for this key (ip)
	ipTokenBucket := t.GetOrCreateIPTokenBucket(req.ip)
	// get the endpointBucket for this key (endpoint)
	endpointBucket := t.GetOrCreateEndpointTokenBucket(req.endpoint)
	return []Limiter{
		userTokenBucket, ipTokenBucket, endpointBucket,
	}
}
