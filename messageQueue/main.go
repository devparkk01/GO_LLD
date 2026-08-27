package main

import (
	"fmt"
	"sync"
)

func runProducer(id int, q *Queue, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		msgID := fmt.Sprintf("prod%d:%d", id, i)
		if err := q.Enqueue(NewMessage(msgID, []byte("Message payload"))); err != nil {
			fmt.Printf("Producer %d: enqueue error: %v\n", id, err)
			return
		}
		fmt.Printf("Producer %d: enqueued %s\n", id, msgID)
	}
}

// runConsumer dequeues messages in a loop until the queue is closed and drained.
func runConsumer(id int, q *Queue, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg := q.Dequeue()
		if msg == nil {
			// we are done consuming
			return
		}
		fmt.Printf("Consumer %d Dequeued %s\n", id, msg.ID)
	}
}

func main() {
	q := NewQueue(5)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Start consumers before producers so they are ready to receive
	// messages immediately and don't miss early wakeups.
	for i := 0; i < 3; i++ {
		consumerWg.Add(1)
		go runConsumer(i, q, &consumerWg)
	}
	
	for i := 0; i < 3; i++ {
		producerWg.Add(1)
		go runProducer(i, q, &producerWg)
	}
	producerWg.Wait()
	q.Close()
	consumerWg.Wait()

}
