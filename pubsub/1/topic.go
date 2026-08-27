package main

import (
	"fmt"
	"sync"
)

// Topic represents a named message channel.
// Publishers send messages to a topic; all subscribers of that topic
// receive a copy of every message (broadcast / pub-sub delivery).
//
// Concurrency model:
//   - mu is an RWMutex because publish (reads the subscriber map) happens
//     far more frequently than subscribe / unsubscribe (writes the map).
//   - Multiple publishers can deliver to the same topic concurrently —
//     they all hold the read lock and only contend at the subscriber level.
type Topic struct {
	name        string
	subscribers map[string]*Subscriber // key: subscriberID, value: *Subscriber
	mu          sync.RWMutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		name:        name,
		subscribers: make(map[string]*Subscriber),
	}
}

// addSubscriber creates a new Subscriber for the given ID and registers it
// on this topic. Returns an error if a subscriber with that ID already exists
func (t *Topic) addSubscriber(subscriberId string, handler func(*Message)) (*Subscriber, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	// check if subscriber exists already 
	if _, ok := t.subscribers[subscriberId]; ok {
		return nil, fmt.Errorf("Subscriber %s already exists \n", subscriberId)
	}
	// Create a new subscriber 
	s := NewSubscriber(subscriberId, SUBSCRIBER_CAPACITY, handler)
	t.subscribers[subscriberId] = s
	return s, nil
}

// removeSubscriber closes the subscriber's goroutine and removes it from
// the topic.
func (t *Topic) removeSubscriber(subscriberId string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	// check if subscriber exists
	sub, ok := t.subscribers[subscriberId]
	if !ok {
		return fmt.Errorf("Subscriber %s has not subscribed to topic %s", subscriberId, t.name)
	}

	// Close first — wakes the Start() goroutine so it can exit.
	// Deleting from the map without closing would leave the goroutine
	// sleeping on cond.Wait() forever (goroutine leak).
	sub.Close()
	delete(t.subscribers, subscriberId)
	return nil
}

// publish delivers a copy of the message to every subscriber's inbox.
func (t *Topic) publish(message *Message) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, sub := range t.subscribers {
		sub.deliver(message)
	}
}

// CloseAll closes every subscriber on this topic.
// Called during broker shutdown to unblock all Start() goroutines.
func (t *Topic) CloseAll() {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, sub := range t.subscribers {
		sub.Close()
	}
}
