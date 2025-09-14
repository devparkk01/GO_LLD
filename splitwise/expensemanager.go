package main

import (
	"fmt"
	"sync"
)

type ExpenseManager struct {
	users  map[string]*User // map[userId] = &User 
	groups map[string]*Group  // map[groupId] = &Group
	mu     sync.RWMutex
}

var manager *ExpenseManager

func NewExpenseManager() *ExpenseManager {
	if manager == nil {
		manager = &ExpenseManager{
			users: make(map[string]*User),
			groups: make(map[string]*Group),
		}
	}
	return manager 
}

func(t *ExpenseManager) CreateUser(id, name, email string) *User {
	t.mu.Lock() 
	defer t.mu.Unlock()

	user := NewUser(id, name, email)
	t.users[user.id] = user 
	return user 
}

func (t *ExpenseManager) CreateGroup(id, name string) *Group {
	t.mu.Lock() 
	defer t.mu.Unlock()

	group := NewGroup(id, name)
	t.groups[group.id] = group 
	return group 
}

func (t *ExpenseManager) AddMemberToGroup(groupId string, userId string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	group, exists := t.groups[groupId]
	if !exists {
		fmt.Printf("Groupid %s does not exist\n", groupId)
		return
	}
	user, exists := t.users[userId]
	if !exists{
		fmt.Printf("Userid %s does not exist\n", userId)
		return 
	}
	group.AddMember(user)
	
}


func (t *ExpenseManager) CreateExpense(id string, description string, payer *User, amountPaid float64, participants []*User, splitStrategy SplitStrategy) *Expense {
	return NewExpense(id, description, payer, amountPaid, participants, splitStrategy)
}


func (e *ExpenseManager) AddExpenseToGroup(groupId string, expense *Expense) {
	e.mu.RLock() // acquire the reader lock
	group, exists := e.groups[groupId]
	e.mu.RUnlock() // release the reader lock first 

	if !exists {
		fmt.Println("This group does not exist")
		return 
	}

	// We can have another validation here to check if the participants are part of the group or not. 

	// calculate the splits 
	splits, err := expense.splitStrategy.CalculateSplits(expense)
	if err != nil {
		fmt.Println("Unable to split the bill. please check the splits")
		return 
	}
	payer := expense.payer

	// update balances using the splits 
	for _, split := range splits{
		user := split.user
		// skip if the user is same as payer 
		if user.id == payer.id {
			continue 
		}
		// update balance sheet
		payer.UpdateBalanceSheet(user.id, split.amount)
		user.UpdateBalanceSheet(payer.id, -split.amount)
	}
	// update splits of this expense 
	expense.SetSplits(splits)
	// Add this expense to the group 
	group.AddExpense(expense)
	fmt.Printf("Expense '%s' of %.2f added to group '%s'\n", expense.description, expense.amountPaid, group.name)
}

func(e *ExpenseManager) SettleBalance(userId, otherUserId string) {
	e.mu.RLock()
	user, ok1  := e.users[userId]
	otherUser, ok2  := e.users[otherUserId]
	e.mu.RUnlock()

	if !ok1 {
		fmt.Println("user does not exist", userId)
	}
	if !ok2 {
		fmt.Println("user does not exist ", otherUserId)
	}

	user.SettleBalance(otherUserId)
	otherUser.SettleBalance(userId)

}

func(e *ExpenseManager) TotalOwed(user *User) float64 {
	return user.TotalOwed()
}

func(e *ExpenseManager) TotalIsOwed(user *User) float64 {
	return user.TotalIsOwed()
}