package main

import "sync"

type Topic struct {
	name        string
	subscribers map[Subscriber]struct{} // set of subscribers subscribed to this topic 
	mu          sync.RWMutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		name: name, 
		subscribers: make(map[Subscriber]struct{}),
	}
}

// Adds a subscriber to the topic 
func (t *Topic) AddSubscriber(s Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.subscribers[s] = struct{}{}
}


// Removes a subscriber from the topic 
func (t *Topic) RemoveSubscriber(s Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.subscribers, s)
}

// Publishes a message to the all subscribers 
func (t *Topic) Publish(m *Message) {
	t.mu.RLock()
	defer t.mu.RUnlock() 
	for s := range t.subscribers {
		s.onMessage(m)
	}
}