package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	TRANSACTION_STATUS_SUCCESS TransactionStatus = "Success"
	TRANSACTION_STATUS_FAILED  TransactionStatus = "Failed"
)

type Transaction struct {
	transactionId string
	fromWalletId  string
	toWalletId    string
	amount        float64
	fromCurrency  Currency
	toCurrency    Currency
	timestamp     time.Time
	status        TransactionStatus
}

func NewTransaction(fromWalletId, toWalletId string, fromCurrency, toCurrency Currency, amount float64, status TransactionStatus) *Transaction {
	return &Transaction{
		transactionId: "t-" + uuid.NewString(),
		fromWalletId:  fromWalletId,
		toWalletId:    toWalletId,
		fromCurrency:  fromCurrency,
		toCurrency:    toCurrency,
		amount:        amount,
		timestamp:     time.Now(),
		status:        status,
	}
}

func DisplayTransactions(transactions []*Transaction) {
	for _, t := range transactions {
		fmt.Printf("transactionID: %s, fromWallet: %s, toWallet: %s, fromCurrency:%s, toCurrency: %s, amount: %.2f, status: %s \n", t.transactionId, t.fromWalletId, t.toWalletId, t.fromCurrency.GetCode(), t.toCurrency.GetCode(), t.amount ,t.status)
	}
}
