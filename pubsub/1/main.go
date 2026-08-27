package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	publishWg := sync.WaitGroup{}
	broker := NewBroker()
	broker.CreateTopic("Events")
	broker.CreateTopic("Logs")
	broker.CreateTopic("Notifications")

	broker.Subscribe("Events", "sub1", func(m *Message) {
		fmt.Printf("[sub1][Events] %s, %s\n", m.Id, m.Payload)
	})
	broker.Subscribe("Events", "sub2", func(m *Message) {
		fmt.Printf("[sub2][Events] %s, %s\n", m.Id, m.Payload)
	})
	broker.Subscribe("Logs", "sub1", func(m *Message) {
		fmt.Printf("[sub1][Logs] %s, %s\n", m.Id, m.Payload)
	})

	broker.Subscribe("Notifications", "sub2", func(m *Message) {
		fmt.Printf("[sub2][Notifications] %s\n", m.Payload)
	})

	messages := []*Message{
		NewMessage("m1", "user signed up", "Events", time.Now()),
		NewMessage("m2", "payment failed", "Events", time.Now()),
		NewMessage("m3", "server error", "Notifications", time.Now()),
		NewMessage("m4", "Payment accepted", "Events", time.Now()),
		NewMessage("m5", "User created", "Logs", time.Now()),
		NewMessage("m6", "server error", "Logs", time.Now()),
	}


	for _, msg := range messages {
		publishWg.Add(1)
		go func() {
			defer publishWg.Done() 
			broker.Publish(msg, msg.Topic)
		}()
	}

	publishWg.Wait()

	broker.Close()
	broker.Wait()

}
