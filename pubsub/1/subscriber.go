package main

import (
	"container/list"
	"sync"
)

const (
	SUBSCRIBER_CAPACITY int = 10
)

// Subscriber represents a single consumer of a topic.
// Each subscriber has its own bounded inbox — a FIFO queue backed by a
// linked list — and runs a dedicated goroutine (Start) that drains the
// inbox by calling the handler for each message.
//
// Concurrency model:
//   - mu protects all fields (inbox, closed)
//   - cond is used by Start() to sleep efficiently when the inbox is empty,
//     and by deliver() / Close() to wake it up
//   - deliver() and Start() never hold the lock at the same time for long;
//     the handler is called outside the lock to avoid blocking producers
type Subscriber struct {
	id       string
	inbox    *list.List
	capacity int
	handler  func(*Message) // caller-provided processing logic
	closed   bool
	mu       sync.Mutex
	cond     *sync.Cond
}

// NewSubscriber creates a Subscriber with a linked-list inbox and
// initialises the condition variable tied to the subscriber's own mutex.
func NewSubscriber(id string, capacity int, handler func(*Message)) *Subscriber {
	s := &Subscriber{
		id:       id,
		inbox:    list.New(),
		capacity: capacity,
		handler:  handler,
	}
	s.cond = sync.NewCond(&s.mu)
	return s
}

// deliver puts a message into the subscriber's inbox.
// It is called by the topic's publish path and must return quickly.
//
// Drop policy: DropNewest — if the inbox is full, the incoming message
// is silently discarded and the existing messages are preserved in order.
// This ensures the publisher never blocks on a slow subscriber.
func (s *Subscriber) deliver(msg *Message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Do not accept new messages if the subscriber is closed or inbox is full.
	if s.closed || s.inbox.Len() >= s.capacity {
		return
	}

	s.inbox.PushBack(msg)
	// Signal wakes exactly one goroutine waiting in Start().
	// Broadcast is not needed here — one message was added, one consumer should wake.
	s.cond.Signal()
}

// Start is the subscriber's consume loop. It blocks when the inbox is empty
// and processes one message at a time by calling the handler.
//
// Important: the handler is called outside the lock. This prevents a slow
// handler from blocking deliver() and by extension the entire publish path.
//
// Start exits cleanly when Close() is called and the inbox is empty.
func (s *Subscriber) Start() {
	for {
		s.mu.Lock()
		// Wait until there is a message or the subscriber is closed.
		// This must be a for loop — not an if — because:
		//   1. Spurious wakeups: Wait() can return without a Signal.
		//   2. After waking, another goroutine may have already consumed the message.
		for s.inbox.Len() == 0 {
			if s.closed {
				s.mu.Unlock()
				return
			}
			s.cond.Wait() // atomically releases mu and sleeps until signalled
		}

		front := s.inbox.Front()
		// Dequeue the oldest message (FIFO).
		s.inbox.Remove(front)
		msg := front.Value.(*Message)

		// Release the lock before calling the handler.
		// Holding the lock during handler execution would block deliver()
		// for the entire duration of message processing.
		s.mu.Unlock()
		s.handler(msg)
	}
}

// Close marks the subscriber as closed and wakes all goroutines blocked
// in Start() so they can check the closed flag and exit.
// Broadcast is used here (not Signal) because there may be multiple
// goroutines waiting — all of them should exit.
func (s *Subscriber) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closed = true
	s.cond.Broadcast() // wake all blocked goroutine
}
