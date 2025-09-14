package main

import (
	"fmt"
	"sync"
)

type Group struct {
	id       string
	name     string
	members  []*User
	expenses []*Expense
	mu       sync.Mutex
}

func NewGroup(id string, name string) *Group {
	return &Group{
		id:   id,
		name: name,
	}
}

func (g *Group) AddExpense(expense *Expense) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.expenses = append(g.expenses, expense)
}

func(g *Group) AddMember(user *User) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.members = append(g.members, user)
	fmt.Printf("Added user %s to group %s \n", user.id, g.id)
}