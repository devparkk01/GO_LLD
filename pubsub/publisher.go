package main

import (
	"fmt"
	"sync"
)

type Publisher struct {
	topics map[*Topic]struct{} // Set of topics
	mu sync.RWMutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		topics: make(map[*Topic]struct{}),
	}
}


// Registers a topic to the Publisher 
func (p *Publisher) RegisterTopic(t *Topic) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.topics[t] = struct{}{}
}

// Publishes a message to the topic 
func (p *Publisher) Publish(t *Topic, m *Message) {
	// get the topic
	p.mu.RLock()
	_, ok := p.topics[t]
	p.mu.RUnlock()
	if !ok {
		fmt.Printf("This publisher can't publish to topic %s\n", t.name)
		return 
	}
	t.Publish(m) 
}
