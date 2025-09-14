package main

import "sync"

type Product struct {
	id           string
	name         string
	description  string
	categoryId   string
	pricePerItem float32
	quantity     int
	mu           sync.RWMutex
	// attributes []string
}

func(p *Product) GetQuantity() int {
	p.mu.RLock()
	defer p.mu.RUnlock() 
	return p.quantity 
}

func(p *Product) UpdateQuantity(quantity int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.quantity += quantity 
}

func(p *Product) IsAvailable(quantity int) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.quantity >= quantity 
}