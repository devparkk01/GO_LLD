# Point-to-Point Message Queue (In-Memory)

A bounded, thread-safe, in-memory point-to-point message queue in Go.  
Multiple producers and consumers operate concurrently. Each message is delivered to **exactly one consumer**.

---

## Design Goals


| Goal                              | Decision                                    |
| --------------------------------- | ------------------------------------------- |
| Exactly-once delivery per message | Competing consumers via shared queue        |
| Bounded memory                    | Fixed-capacity ring buffer                  |
| No busy-waiting                   | `sync.Cond` for blocking and wakeup         |
| Drain on shutdown                 | Consumers dequeue until empty, then exit    |
| Thread safety                     | Single `sync.Mutex` guards all shared state |


---

## Architecture

```
Producer 1 ─┐
Producer 2 ─┼──▶  [ Queue: ring buffer ] ──▶ one of the consumers gets each message
Producer 3 ─┘         (bounded, FIFO)
                                         ├──▶ Consumer 1
                                         ├──▶ Consumer 2
                                         └──▶ Consumer 3
```

---

## Core Entities

### `Message`

The unit of data in the system.

```go
type Message struct {
    ID        string
    Payload   []byte
    CreatedAt time.Time
}
```

No topic or routing metadata — this is point-to-point, not pub-sub.

### `Queue`

The central entity. Owns the buffer and all synchronisation primitives.

```go
type Queue struct {
    buffer   []Message    // ring buffer
    capacity int
    head     int          // next read index
    tail     int          // next write index
    count    int          // current message count

    mu       sync.Mutex
    notEmpty *sync.Cond   // consumers wait here
    notFull  *sync.Cond   // producers wait here

    closed   bool
}
```

No separate `Producer` or `Consumer` structs — any goroutine calling `Enqueue` is a producer, any goroutine calling `Dequeue` is a consumer.

---

## Internal Data Structure — Ring Buffer

A ring buffer is a fixed-size array where the end wraps back to the beginning using modulo arithmetic. No memory allocation after initialisation, O(1) enqueue and dequeue.

```
capacity = 5, after enqueuing A B C:

index:  0    1    2    3    4
      [ A  | B  | C  |    |   ]
        ↑              ↑
       head           tail

After dequeuing A, enqueuing D E F (tail wraps):

index:  0    1    2    3    4
      [ F  | B  | C  | D  | E ]
        ↑    ↑
       tail head
```

Wrap formula:

```go
tail = (tail + 1) % capacity
head = (head + 1) % capacity
```

### Why not other options?


| Option           | Enqueue | Dequeue | Notes                                                      |
| ---------------- | ------- | ------- | ---------------------------------------------------------- |
| **Ring buffer**  | O(1)    | O(1)    | Fixed memory, cache friendly — best for bounded queues     |
| Slice + shift    | O(1)    | O(n)    | Shifts all elements on dequeue — never use for queues      |
| Linked list      | O(1)    | O(1)    | Per-node heap allocation, poor cache locality, GC pressure |
| Buffered channel | O(1)    | O(1)    | Loses fine-grained control over errors and shutdown        |


---

## Synchronisation Design

### One mutex, two `sync.Cond`

All shared state (`buffer`, `head`, `tail`, `count`, `closed`) is guarded by a single `sync.Mutex`. Two conds share this mutex:


| Cond       | Who waits                  | Who signals               |
| ---------- | -------------------------- | ------------------------- |
| `notEmpty` | Consumers (queue is empty) | Producers (after enqueue) |
| `notFull`  | Producers (queue is full)  | Consumers (after dequeue) |


**Why two conds?**  
With one cond, `Signal()` wakes an arbitrary waiter — it might wake another producer instead of a waiting consumer. You'd need `Broadcast` everywhere, causing a thundering herd on every operation.

Two conds allow precise wakeup: `notEmpty.Signal()` always wakes a consumer, `notFull.Signal()` always wakes a producer.

**Why one mutex?**  
All conds protect the same shared state. Two mutexes would allow concurrent modification of `count` by a producer and consumer simultaneously — a data race.

### Signal vs Broadcast


| Situation        | Call                                           | Reason                                |
| ---------------- | ---------------------------------------------- | ------------------------------------- |
| Message enqueued | `notEmpty.Signal()`                            | One message added → wake one consumer |
| Message dequeued | `notFull.Signal()`                             | One slot freed → wake one producer    |
| Queue closed     | `notEmpty.Broadcast()` + `notFull.Broadcast()` | Wake all blocked goroutines to exit   |


### Spurious wakeups

`sync.Cond.Wait()` can return without being signalled. Always re-check the condition in a `for` loop, never `if`:

```go
// correct
for q.count == 0 && !q.closed {
    q.notEmpty.Wait()
}

// wrong — misses spurious wakeups
if q.count == 0 && !q.closed {
    q.notEmpty.Wait()
}
```

---

## API

```go
func NewQueue(capacity int) *Queue
func (q *Queue) Enqueue(msg Message) error   // blocks when full; error if closed
func (q *Queue) Dequeue() *Message           // blocks when empty; nil when closed+drained
func (q *Queue) Close()                      // idempotent shutdown
```

---

## Shutdown — Drain on Close

```
Close() called
    │
    ├── sets closed = true
    ├── Broadcast on notFull  → unblocks all producers → they return error
    └── Broadcast on notEmpty → unblocks all consumers → they re-evaluate

Consumer wakes up:
    ├── count > 0 → dequeue message, continue loop
    └── count == 0 → return nil → caller exits
```

Producers stop immediately on close. Consumers drain all remaining messages before exiting.

---

## Running

```bash
go run --race .
```

Expected output (order will vary):

```
Producer 1: enqueued prod1:0
Consumer 2: dequeued prod1:0
Producer 2: enqueued prod2:0
Consumer 1: dequeued prod2:0
...
Consumer 3: queue drained, exiting
shutdown complete — all messages drained
```

---

## Out of Scope

- Message acknowledgement / redelivery on consumer crash
- Persistence / durability
- Priority queues
- Dead letter queues
- Per-message TTL

