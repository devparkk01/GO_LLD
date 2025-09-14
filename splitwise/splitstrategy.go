package main

import (
	"fmt"
	"math"
)

type SplitStrategy interface {
	CalculateSplits(expense *Expense) ([]*Split, error)
}

type EqualSplit struct{}

func (es *EqualSplit) CalculateSplits(expense *Expense) ([]*Split, error) {
	n := len(expense.participants)
	if n == 0 {
		return nil, fmt.Errorf("no participants for expense")
	}
	// calculate per head amount 
	perHead := expense.amountPaid / float64(n)
	splits := make([]*Split, 0, n)

	for _, user := range expense.participants {
		splits = append(splits, NewSplit(user, perHead))
	}
	return splits, nil
}

type PercentageSplit struct {
	Percentages map[string]float64  // key: userID, value: percentage share
}

func (ps *PercentageSplit) CalculateSplits(expense *Expense) ([]*Split, error) {
	totalPercentage := 0.0
	// Add all percententages 
	for _, percent := range ps.Percentages {
		totalPercentage += percent
	}
	// if difference is more than 0.01 then return error 
	if math.Abs(totalPercentage-100.0) > 0.01 {
		return nil, fmt.Errorf("percentages do not sum to 100")
	}

	splits := make([]*Split, 0, len(expense.participants))
	for _, user := range expense.participants {
		percent, ok := ps.Percentages[user.id]
		if !ok {
			return nil, fmt.Errorf("no percent specified for user %s", user.id)
		}
		amount := (percent / 100.0) * expense.amountPaid
		splits = append(splits, NewSplit(user, amount))
	}

	return splits, nil
}

type UnequalSplit struct {
	SplitMap map[string]float64  // key: userId, value: exact amount 
}

func (us *UnequalSplit) CalculateSplits(expense *Expense) ([]*Split, error) {
	total := 0.0 // should be equal to total amount, which is expense.AmountPaid
	splits := make([]*Split, 0, len(expense.participants))

	for _, user := range expense.participants {
		amount, ok := us.SplitMap[user.id]
		if !ok {
			return nil, fmt.Errorf("no split specified for user %s", user.id)
		}
		splits = append(splits, NewSplit(user, amount))
		total += amount
	}

	if math.Abs(total-expense.amountPaid) > 0.01 {
		return nil, fmt.Errorf("sum of splits %.2f does not match expense amount %.2f", total, expense.amountPaid)
	}

	return splits, nil
}