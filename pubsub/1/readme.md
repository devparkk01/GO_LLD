# In-Memory Pub-Sub Message Broker

A thread-safe, in-memory publish-subscribe message broker implemented in Go.
Designed as an LLD interview exercise demonstrating goroutines, channels, `sync.Mutex`, `sync.RWMutex`, and `sync.Cond`.

---

## Design

### Delivery model

**Pub-Sub (broadcast)** — every message published to a topic is delivered to all subscribers of that topic. Each subscriber receives its own independent copy.

```
Publisher
    │
    ▼
  Broker
    │
    ├──▶ Topic: "Events"
    │        ├──▶ Subscriber "sub1" inbox: [msg1, msg2]
    │        └──▶ Subscriber "sub2" inbox: [msg1, msg2]
    │
    └──▶ Topic: "Logs"
             └──▶ Subscriber "sub1" inbox: [msg3]
```

### Components

| Component | Responsibility |
|---|---|
| `Message` | Data unit — carries ID, payload, topic name, and timestamp |
| `Subscriber` | Owns a bounded inbox and a goroutine that processes messages via a handler callback |
| `Topic` | Named channel — owns a map of subscribers, fans out published messages to all of them |
| `Broker` | Entry point — manages topics, starts subscriber goroutines, coordinates shutdown |

---

## Concurrency design

### Locks

| Lock | Where | Why |
|---|---|---|
| `broker.mu` (RWMutex) | Broker | Publish and Subscribe read the topic map concurrently; CreateTopic writes it rarely |
| `topic.mu` (RWMutex) | Topic | publish() reads the subscriber map; addSubscriber/removeSubscriber write it rarely |
| `subscriber.mu` (Mutex) | Subscriber | Protects inbox and closed flag accessed by deliver() and Start() |

### sync.Cond

Each subscriber uses a `sync.Cond` tied to its own mutex to coordinate between the producer (`deliver`) and consumer (`Start`) sides:

- `deliver()` calls `cond.Signal()` after pushing a message — wakes exactly one waiting `Start()` goroutine
- `Start()` calls `cond.Wait()` when the inbox is empty — sleeps with zero CPU until signalled
- `Close()` calls `cond.Broadcast()` — wakes all goroutines so they can check the closed flag and exit

### Drop policy — DropNewest

When a subscriber's inbox is full, the incoming message is silently dropped and the publisher returns immediately. This ensures:

- The publisher **never blocks** on a slow subscriber
- Existing messages in the inbox are **preserved in order**
- The system stays fully **decoupled** — slow consumers do not stall producers

### Message flow

```
broker.Publish(msg, "Events")
    └── topic.publish(msg)          [holds topic.mu read lock]
            └── sub.deliver(msg)    [holds sub.mu — O(1) list append]
                    └── cond.Signal()

                            ▼  (async, in subscriber's goroutine)

                    sub.Start() wakes from cond.Wait()
                            └── dequeues message
                            └── releases sub.mu
                            └── calls handler(msg)  [lock-free]
```

---

## Shutdown sequence

```
broker.Close()
    └── sets broker.closed = true
    └── topic.CloseAll()
            └── sub.Close()  for each subscriber
                    └── sets sub.closed = true
                    └── cond.Broadcast()  — wakes Start() goroutine

Start() goroutine wakes
    └── inbox is empty AND closed == true
    └── returns

broker.wg.Done() fires (deferred in Subscribe goroutine wrapper)

broker.Wait()
    └── blocks until broker.wg counter reaches zero
    └── returns — all goroutines have exited cleanly
```

**Rule:** always publish all messages before calling `Close()`. Since `deliver()` is synchronous and O(1), all messages are guaranteed to be in subscriber inboxes before the publish loop returns. It is safe to call `Close()` immediately after the last `Publish`.

---

## API

```go
// Broker
broker := NewBroker()
broker.CreateTopic("Events")                          // register a topic
broker.Subscribe("Events", "sub1", func(m *Message) { // register handler, start goroutine
    fmt.Println(m.Payload)
})
broker.Unsubscribe("Events", "sub1")                  // stop goroutine, remove from topic
broker.Publish(msg, "Events")                         // deliver to all subscribers
broker.Close()                                        // signal all goroutines to exit
broker.Wait()                                         // block until all goroutines exit

// Message
msg := NewMessage("m1", "user signed up", "Events", time.Now())
```

---

## File structure

```
message.go      — Message struct and constructor
subscriber.go   — Subscriber struct, deliver(), Start(), Close()
topic.go        — Topic struct, addSubscriber(), removeSubscriber(), publish(), CloseAll()
broker.go       — Broker struct, full public API
main.go         — Example wiring and usage
```

---

## Key design decisions

**Why `sync.Cond` instead of a channel for the inbox?**

A channel-based inbox would work for simple cases but `sync.Cond` is more appropriate here because the wait condition involves inspecting struct state (`inbox.Len() == 0`, `closed`) that is already protected by a mutex. `sync.Cond` keeps the data and the synchronisation together without transferring ownership.

**Why `container/list` instead of a slice?**

A linked list gives O(1) push to back and O(1) pop from front — the two operations a FIFO queue needs. A naive slice has O(n) pop from front due to element shifting. A ring buffer (circular slice) would also give O(1) for both with better cache locality, but adds implementation complexity.

**Why `RWMutex` on Broker and Topic?**

Publish reads the topic/subscriber maps far more frequently than Subscribe writes them. `RWMutex` allows unlimited concurrent readers — multiple publishers can deliver to the same topic simultaneously. Writes (CreateTopic, Subscribe, Unsubscribe) are exclusive but rare.

**Why is the handler called outside the subscriber's lock?**

Holding `sub.mu` during handler execution would block `deliver()` for the entire duration of message processing. Slow handlers would stall the publish path. Releasing the lock before calling the handler keeps deliver() fast and the system decoupled.

**Why DropNewest over DropOldest or Block?**

- **Block** — publisher waits for slow subscriber — couples producer and consumer, defeating async design
- **DropOldest** — keeps latest messages — good for live dashboards, but requires an extra pop+push and disturbs existing inbox order
- **DropNewest** — discards incoming message — simplest implementation, preserves existing message order, publisher never blocks

---

## Limitations and possible extensions

| Limitation | Extension |
|---|---|
| No message persistence | Add a message log with per-subscriber offset (Kafka-style) |
| No acknowledgement | Add Ack() — redeliver on timeout |
| No metrics | Track drop count, queue depth, processing latency per subscriber |
| No priority | Add a priority field to Message and use a heap instead of a list |
| Subscriber tied to one topic | Allow one subscriber object across multiple topics with a shared inbox |
| No dead letter queue | Route failed/dropped messages to a DLQ topic |