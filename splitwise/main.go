package main

import "fmt"

func main() {
	manager := NewExpenseManager()

	alice := manager.CreateUser("u1", "Alice", "alice@test.com")
	bob := manager.CreateUser("u2", "Bob", "bob@test.com")
	charlie := manager.CreateUser("u3", "Charlie", "charlie@test.com")

	friendsGroup := manager.CreateGroup("g1", "Friends")

	manager.AddMemberToGroup(friendsGroup.id, alice.id)
	manager.AddMemberToGroup(friendsGroup.id, bob.id)
	manager.AddMemberToGroup(friendsGroup.id, charlie.id)

	expense := manager.CreateExpense("e1", "Dinner", alice, 90, []*User{alice, bob, charlie}, &EqualSplit{})

	manager.AddExpenseToGroup(friendsGroup.id, expense)

	fmt.Println(manager.TotalIsOwed(alice))
	fmt.Println(manager.TotalIsOwed(bob))
	fmt.Println(manager.TotalIsOwed(charlie))

	fmt.Println(manager.TotalOwed(alice))
	fmt.Println(manager.TotalOwed(bob))
	fmt.Println(manager.TotalOwed(charlie))

	manager.SettleBalance(alice.id, bob.id )
	manager.SettleBalance(alice.id, charlie.id)

	fmt.Println(manager.TotalIsOwed(alice))
	fmt.Println(manager.TotalIsOwed(bob))
	fmt.Println(manager.TotalIsOwed(charlie))

	fmt.Println(manager.TotalOwed(alice))
	fmt.Println(manager.TotalOwed(bob))
	fmt.Println(manager.TotalOwed(charlie))

}
