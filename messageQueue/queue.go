package main

import (
	"fmt"
	"sync"
)

// Queue is a bounded, thread-safe, point-to-point message queue.
// Multiple producers and consumers can operate concurrently.
// Each message is delivered to exactly one consumer.
//
// Internally uses a ring buffer for O(1) enqueue and dequeue,
// and two sync.Cond variables to coordinate blocking and wakeup
// between producers and consumers without busy-waiting.
type Queue struct {
	buffer   []Message // fixed-size ring buffer
	capacity int       // maximum number of messages the queue can hold
	head     int       // index of the next message to be dequeued
	tail     int       // index where the next message will be enqueued
	count    int       // current number of messages in the queue

	mu       sync.Mutex // guards all shared state
	notEmpty *sync.Cond // consumers wait here when queue is empty
	notFull  *sync.Cond // producers wait here when queue is full

	closed bool // true after Close() is called; no new enqueues allowed
}

func NewQueue(capacity int) *Queue {
	// Pre-allocate the full buffer so index writes work immediately.
	// length == capacity (not 0) is required for ring buffer index access.
	q := &Queue{
		buffer:   make([]Message, capacity),
		capacity: capacity,
	}
	// Both conds share the same mutex because they protect the same state.
	// Two separate conds allow precise wakeup:
	//   notEmpty.Signal() wakes a consumer only
	//   notFull.Signal()  wakes a producer only
	// Using one cond would require Broadcast everywhere, causing thundering herd.
	q.notEmpty = sync.NewCond(&q.mu)
	q.notFull = sync.NewCond(&q.mu)
	return q
}

func (q *Queue) Enqueue(msg Message) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Wait while q is full and q is not closed
	// Loop (not if) guards against spurious wakeups — always re-check
	// the condition after Wait() returns.
	for q.count == q.capacity && !q.closed {
		q.notFull.Wait()
	}

	// After waking, reject the message if the queue was closed.
	if q.closed {
		return fmt.Errorf("Enqueue on closed queue")
	}
	// Write to tail position and advance tail using modulo wrap.
	// This is what makes it a ring buffer — tail loops back to 0
	// after reaching the end of the slice.
	q.buffer[q.tail] = msg
	q.tail = (q.tail + 1) % q.capacity
	q.count++
	// Wake exactly one blocked consumer. Signal (not Broadcast) is correct
	// here — only one message was added, so only one consumer should wake.
	q.notEmpty.Signal()
	return nil
}

func (q *Queue) Dequeue() *Message {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Wait while q is empty and q is not closed
	// The loop handles spurious wakeups — condition must be re-checked
	// each time Wait() returns.
	for q.count == 0 && !q.closed {
		q.notEmpty.Wait()
	}

	// If we woke because of Close() and there are no messages left,
	// signal drain completion to the caller.
	// Important: if closed == true but count > 0, we still dequeue —
	// this is the drain-on-shutdown behaviour.
	if q.count == 0 {
		return nil // closed and drained
	}
	msg := q.buffer[q.head]
	// Zero out the slot to release the reference to the payload bytes.
	// Without this, the ring buffer holds a pointer to []byte even after
	// the message is consumed, preventing GC from collecting it.
	q.buffer[q.head] = Message{} // GC help
	// Advance head using modulo wrap — mirrors how tail advances in Enqueue.
	q.head = (q.head + 1) % q.capacity
	q.count--
	// Wake exactly one blocked producer — one slot just opened up.
	q.notFull.Signal()
	return &msg

}

// Close shuts down the queue. After Close:
//   - Enqueue returns an error immediately.
//   - Dequeue continues to drain remaining messages.
//   - Dequeue returns nil once the queue is empty.
//
// Close is idempotent — safe to call multiple times.
func (q *Queue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return
	}
	q.closed = true
	q.notEmpty.Broadcast() // unblocks all waiting consumers
	q.notFull.Broadcast()  // unblocks all waiting producers
}
