package main

import "fmt"

func main() {
	usd := NewCurrency("USD")
	inr := NewCurrency("INR")
	eur := NewCurrency("EUR")

	bankTransfer := &BankTransfer{}
	ccs := &CurrencyConversionService{
		rates: map[string]map[string]float64{
			usd.GetCode(): {inr.GetCode(): 80.0, eur.GetCode(): 0.9},
			inr.GetCode(): {usd.GetCode(): 0.0125, eur.GetCode(): 0.011},
			eur.GetCode(): {usd.GetCode(): 1.1, inr.GetCode(): 90.0},
		},
	}

	ws := NewWalletService(ccs)

	alice := ws.CreateUser("u1", "alice")
	bob := ws.CreateUser("u2", "bob")
	charlie := ws.CreateUser("u3", "charlie")
	david := ws.CreateUser("u4", "david")

	// load balance
	err := ws.LoadBalance(alice.GetId(), usd, 50, bankTransfer)
	if err != nil {
		fmt.Println("Failed to load balance.", err.Error())
		return
	}
	_ = ws.GetAllBalance(alice.GetId())

	err = ws.LoadBalance(bob.GetId(), usd, 20, bankTransfer)
	if err != nil {
		fmt.Println("Failed to load balance.", err.Error())
		return
	}
	_ = ws.GetAllBalance(bob.GetId())

	err = ws.LoadBalance(charlie.GetId(), eur, 60, bankTransfer)
	if err != nil {
		fmt.Println("Failed to load balance.", err.Error())
		return
	}
	_ = ws.GetAllBalance(charlie.GetId())

	err = ws.LoadBalance(david.GetId(), inr, 2600, bankTransfer)
	if err != nil {
		fmt.Println("Failed to load balance.", err.Error())
		return
	}
	_ = ws.GetAllBalance(david.GetId())

	// Transfer
	err = ws.Transfer(alice.GetId(), david.GetId(), usd, usd, 20)
	if err != nil {
		fmt.Println("Failed to transfer ", err.Error())
		return
	}

	_ = ws.GetAllBalance(alice.GetId())
	_ = ws.GetAllBalance(david.GetId())

	// Transfer ( 2000 INR from david to charlie in USD)
	err = ws.Transfer(david.GetId(), charlie.GetId(), inr, usd, 2000)
	if err != nil {
		fmt.Println("Failed to transfer ", err.Error())
		return
	}

	_ = ws.GetAllBalance(david.GetId())
	_ = ws.GetAllBalance(charlie.GetId())

	davidTransactions, _ := ws.GetTransactions(david.GetId())
	charlieTransactions, _ := ws.GetTransactions(charlie.GetId())
	fmt.Println("transactions for David")
	DisplayTransactions(davidTransactions)
	fmt.Println("transactions for charlie")
	DisplayTransactions(charlieTransactions)

}
