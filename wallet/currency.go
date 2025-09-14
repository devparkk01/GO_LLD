package main 

type Currency struct {
	code string // e.g., "USD", "EUR", "INR"
}

func NewCurrency(code string) Currency {
	return Currency{code: code }
}

func (c *Currency) GetCode() string {
	return c.code 
}

