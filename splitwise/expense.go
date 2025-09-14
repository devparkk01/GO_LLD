package main 


type Expense struct {
	id string 
	description string 
	payer *User 
	amountPaid float64
	participants []*User 
	splitStrategy SplitStrategy 
	splits []*Split 
}

func NewExpense(id string, description string, payer *User, amountPaid float64, participants []*User, splitStrategy SplitStrategy) *Expense{
	return &Expense{
		id: id, 
		description: description,
		payer: payer,
		amountPaid: amountPaid,
		participants: participants,
		splitStrategy: splitStrategy,
	}
}

func(e *Expense) SetSplits(splits []*Split) {
	e.splits = splits 
}