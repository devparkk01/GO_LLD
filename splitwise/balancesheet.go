package main

import "sync"

type BalanceSheet struct {
	balances map[string]float64 // otherUserid -> netAmount (positive value means otherUser owes currentUser certainAmount )
	mu       sync.RWMutex
}


func NewBalanceSheet() *BalanceSheet{
	return &BalanceSheet{
		balances: make(map[string]float64),
	}
}

func (bs *BalanceSheet) UpdateBalance(otherUserID string, amount float64) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.balances[otherUserID] += amount
}

func (bs *BalanceSheet) GetBalance(otherUserID string) float64 {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	return bs.balances[otherUserID] // if otherUserID does not exist in this map, it will be returned at 0
}

func (bs *BalanceSheet) SettleBalance(otherUserID string) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	delete(bs.balances, otherUserID)
}

func (bs *BalanceSheet) TotalOwed() float64 {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	total := 0.0
	for _, amount := range bs.balances {
		if amount < 0 {
			total += -amount
		}
	}
	return total
}

func (bs *BalanceSheet) TotalIsOwed() float64 {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	total := 0.0
	for _, amount := range bs.balances {
		if amount > 0 {
			total += amount
		}
	}
	return total
}
