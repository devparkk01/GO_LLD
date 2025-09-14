package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	walletId     string
	userId       string
	balance      map[string]float64 // currrencyCode -> amount
	transactions []*Transaction
	mu sync.RWMutex
}

func NewWallet(userId string) *Wallet {
	return &Wallet{
		walletId: "wallet+" + userId,
		userId: userId,
		balance: make(map[string]float64),
	}
}

func(w *Wallet) GetWalletId() string{
	return w.walletId
}

func(w *Wallet) GetUserId() string{
	return w.userId
}

func (w *Wallet) GetTransactions() []*Transaction {
	return w.transactions
}

func (w *Wallet) AddBalance(currency Currency, amount float64)  {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance[currency.GetCode()] += amount
}

func (w *Wallet) WithdrawBalance(currency Currency, amount float64)  {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance[currency.GetCode()] -= amount
}

// return balance of a particular currency 
func(w *Wallet) GetBalance(currency Currency) float64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.balance[currency.GetCode()]
}

// return balance of all currency 
func(w *Wallet) GetAllBalance() {
	w.mu.RLock()
	defer w.mu.RUnlock()
	balance := w.balance
	
	for k,v := range balance {
		fmt.Printf(" %s: %.2f\n", k, v)
	}
}

func (w *Wallet) AddTransaction(t *Transaction)  {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.transactions = append(w.transactions, t)
}
