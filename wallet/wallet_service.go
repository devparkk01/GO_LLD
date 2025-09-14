package main

import (
	"fmt"
	"sync"
)

type WalletService struct {
	users   map[string]*User
	wallets map[string]*Wallet
	ccs     *CurrencyConversionService
	mu sync.RWMutex
}

var walletServiceInstance *WalletService
var walletServiceOnce sync.Once 

func NewWalletService(ccs *CurrencyConversionService) *WalletService {
	walletServiceOnce.Do(func() {
		walletServiceInstance = &WalletService{
			users: make(map[string]*User),
			wallets: make(map[string]*Wallet),
			ccs: ccs, 
		}
	})
	return walletServiceInstance
}



func(ws *WalletService) CreateUser(id, name string) *User {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	u := NewUser(id, name)
	ws.users[u.GetId()] = u 
	ws.wallets[u.GetWallet().GetWalletId()] = u.GetWallet()

	return u 
}


func(ws *WalletService) LoadBalance(userId string, currency Currency, amount float64, ps PaymentStrategy) error {
	// Get the user and corresponding wallet 
	user, exists := ws.users[userId]
	if !exists {
		return fmt.Errorf("user does not exist %s", userId )
	}
	wallet := user.GetWallet()

	// process the payment 
	if !ps.pay(amount, currency) {
		return fmt.Errorf("failed while processing payment ")
	}

	// update the balance
	wallet.AddBalance(currency, amount)
	return nil 
}

func (ws *WalletService) Transfer(fromUserId, toUserId string, fromCurrency, toCurrency Currency, amount float64) error {
	ws.mu.RLock()
	fromUser := ws.users[fromUserId]
	toUser := ws.users[toUserId]
	ws.mu.RUnlock()

	if fromUser == nil || toUser == nil {
		return fmt.Errorf("one of the user does not exist")
	}

	fromWallet := fromUser.GetWallet()
	toWallet := toUser.GetWallet()

	if fromWallet.GetBalance(fromCurrency) < amount {
		return fmt.Errorf("amount is not sufficient")
	}

	convertedAmount := ws.ccs.Convert(fromCurrency, toCurrency, amount)

	ws.mu.Lock()
	defer ws.mu.Unlock()

	// Deduct 
	fromWallet.WithdrawBalance(fromCurrency, amount)
	// Credit 
	toWallet.AddBalance(toCurrency, convertedAmount)
	// Create transaction
	t := NewTransaction(fromWallet.GetWalletId(), toWallet.GetWalletId(), fromCurrency, toCurrency, amount, TRANSACTION_STATUS_SUCCESS)

	// Add to both users 
	fromWallet.AddTransaction(t)
	toWallet.AddTransaction(t)

	return nil 
}

func(ws *WalletService) GetBalance(userId string, currency Currency) (float64, error) {
	// get the user 
	ws.mu.RLock()
	user, exists := ws.users[userId]
	ws.mu.RUnlock()

	if !exists {
		return 0.00 , fmt.Errorf("User does not exist %s", userId)
	}
	return user.GetWallet().GetBalance(currency), nil
}

func(ws *WalletService) GetAllBalance(userId string) error {
	// get the user 
	ws.mu.RLock()
	user, exists := ws.users[userId]
	ws.mu.RUnlock()

	if !exists {
		return fmt.Errorf("User does not exist %s", userId)
	}
	fmt.Println("Getting balance for ", user.name)
	user.GetWallet().GetAllBalance()
	return nil 
}

func(ws *WalletService) GetTransactions(userId string) ([]*Transaction, error ) {
	// get the user
	ws.mu.RLock()
	user, exists := ws.users[userId]
	ws.mu.RUnlock()
	if !exists {
		return nil, fmt.Errorf("User does not exist %s", userId)
	}
	return user.GetWallet().GetTransactions(), nil 
}