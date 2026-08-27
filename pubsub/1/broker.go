package main

import (
	"fmt"
	"sync"
)

// Broker is the central coordinator of the pub-sub system.
// It owns all topics and manages the lifecycle of subscriber goroutines.

// Concurrency model:
//   - mu is an RWMutex because Publish and Subscribe read the topic map
//     far more often than CreateTopic writes it.
//   - wg tracks all subscriber goroutines started by Subscribe so that
//     Wait() can block until every goroutine has exited cleanly.
//   - closed prevents new operations after Close() is called.
type Broker struct {
	topics map[string]*Topic
	mu     sync.RWMutex
	closed bool
	wg sync.WaitGroup   // tracks all active subscriber goroutines
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string]*Topic),
	}
}

// CreateTopic registers a new named topic on the broker.
// Returns an error if the topic already exists.
func (b *Broker) CreateTopic(topicName string) error  {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.topics[topicName]; ok {
		return fmt.Errorf("Topic %s already exists\n", topicName)
	}
	t := NewTopic(topicName)
	b.topics[t.name] = t
	return nil 
}


// Subscribe registers a handler for a topic and starts a dedicated goroutine
// (via Subscriber.Start) that processes incoming messages.
//
// The goroutine is tracked in the broker's WaitGroup so Wait() can block
// until it exits. wg.Add(1) is called before go to avoid a race where
// Wait() could return before the goroutine is even scheduled.
func (b *Broker) Subscribe(topicName string, subscriberID string, handler func(*Message)) error {
	b.mu.RLock()
	topic, ok := b.topics[topicName]
	b.mu.RUnlock()
	if !ok {
		return fmt.Errorf("Topic %s does not exist\n", topicName)
	}
	s, err := topic.addSubscriber(subscriberID, handler)
	if err != nil {
		return err
	}

	// Track the goroutine before launching it — if we called Add inside
	// the goroutine, Wait() could return before Add() is even reached.
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()  // decrements counter when Start() returns
		s.Start()
	}()
	return nil
}

// Unsubscribe removes a subscriber from a topic and stops its goroutine.
func (b *Broker) Unsubscribe(topicName string, subscriberID string) error {
	b.mu.RLock()
	topic, ok := b.topics[topicName]
	b.mu.RUnlock()
	if !ok {
		return fmt.Errorf("Topic %s does not exist\n", topicName)
	}
	return topic.removeSubscriber(subscriberID)
}

// Publish delivers a message to all subscribers of the named topic.
// Multiple goroutines may call Publish concurrently — the broker's read
// lock allows this; contention is pushed down to the per-subscriber mutex.
//
// Publish is synchronous: it returns only after deliver() has been called
// on every subscriber. Since deliver() is O(1), this is fast in practice.
// Messages are guaranteed to be in subscriber inboxes before Publish returns,
// making it safe to call Close() immediately after the last Publish.
func (b *Broker) Publish(message *Message, topicName string) error {
	b.mu.RLock()
	defer b.mu.RUnlock()
	t, ok := b.topics[topicName]
	if !ok {
		return fmt.Errorf("Topic %s does not exist\n", topicName)
	}
	t.publish(message)
	return nil
}

// Close shuts down the broker by marking it closed and signalling all
// subscriber goroutines to exit. After Close, Publish and Subscribe
// return errors. Call Wait() after Close to block until all goroutines exit.
//
// Shutdown order:
//  1. Close() — marks closed, signals all subscriber goroutines
//  2. Wait()  — blocks until every goroutine has exited
func (b *Broker) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.closed = true
	// Cascade close to every topic, which closes every subscriber,
	// which wakes their Start() goroutines via cond.Broadcast().
	for _, topic := range b.topics {
		topic.CloseAll()
	}
}

// Wait blocks until all subscriber goroutines started by Subscribe have exited.
// Always call Close() before Wait() to ensure goroutines have a signal to exit.
func(b *Broker) Wait(){
	b.wg.Wait() 
}