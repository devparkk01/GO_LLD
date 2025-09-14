package main

import (
	"fmt"
)


func main() {
	ts := NewTweetService()

	alice := ts.CreateUser("user1", "alice", "alice@gmail.com")
	bob := ts.CreateUser("user2", "bob", "bob@gmail.com")
	charlie := ts.CreateUser("user3", "charlie", "charlie@gmail.com",)
	david := ts.CreateUser("user4", "david", "david@gmail.com")


	ts.Follow(bob.id, alice.id)
	ts.Follow(charlie.id, alice.id)
	ts.Follow(alice.id, david.id)
	ts.Follow(charlie.id, david.id)
	ts.Follow(bob.id, david.id)

	t1 := ts.PostTweet("t1", alice.id, "alice says go is awesome", []string{"url1"})
	t2 := ts.PostTweet("t2", david.id, "David says go is awesome", []string{"url1"})
	t3 := ts.PostTweet("t3", charlie.id, "charlie says go is awesome", []string{"url1"})
	t4 := ts.PostTweet("t4", bob.id, "what is go", []string{"url1"})

	ts.LikeTweet(bob.GetId(), t1.GetId())
	ts.LikeTweet(bob.GetId(), t3.GetId())
	ts.LikeTweet(david.GetId(), t2.GetId())
	ts.LikeTweet(david.GetId(), t1.GetId())
	ts.LikeTweet(alice.GetId(), t4.GetId())

	ts.AddComment("c1", david.GetId(), t1.GetId(), "You are right, alice")
	ts.AddComment("c2", charlie.GetId(), t1.GetId(), "You are so right, alice")
	ts.AddComment("c3", bob.GetId(), t1.GetId(), "what are you talking about, alice ?")
	ts.AddComment("c4", alice.GetId(), t4.GetId(), "It's a programming language")
	

	feeds := ts.GetNewsFeed(bob.GetId())
	fmt.Println("Printing feed for bob")
	ts.PrintFeed(feeds)

	// Get notification of Alice 
	fmt.Println("Printing notification for alice ")
	for _, notif := range ts.GetUnreadNotification(alice.GetId()) {
		notif.Display()
	}

	fmt.Println("Printing notification for bob")
	for _, notif := range ts.GetUnreadNotification(bob.GetId()) {
		notif.Display()
	}
	
	// clear all notifications for Alice
	ts.MarkAllRead(alice.GetId())
	fmt.Println("Printing notification for alice ")
	for _, notif := range ts.GetUnreadNotification(alice.GetId()) {
		notif.Display()
	}

}