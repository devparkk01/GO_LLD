package main

import "fmt"

type Subscriber interface {
	onMessage(m *Message) // method implemented by all concrete subscribers
}


// Concrete subscribers
type OrderSubscriber struct {
	name string
}
type EmailSubscriber struct {
	name string
}
type PaymentSubscriber struct {
	name string
}

func NewOrderSubscriber(name string) Subscriber {
	return &OrderSubscriber{
		name: name,
	}
}

func NewPaymentSubscriber(name string) Subscriber {
	return &OrderSubscriber{
		name: name,
	}
}

func NewEmailSubscriber(name string) Subscriber {
	return &OrderSubscriber{
		name: name,
	}
}

func (o *OrderSubscriber) onMessage(m *Message) {
	fmt.Printf("Subscriber %s received a message. Message id: %s, payload: %s\n", o.name, m.id, m.payload)
}

func (o *EmailSubscriber) onMessage(m *Message) {
	fmt.Printf("Subscriber %s received a message. Message id: %s, payload: %s\n", o.name, m.id, m.payload)

}

func (o *PaymentSubscriber) onMessage(m *Message) {
	fmt.Printf("Subscriber %s received a message. Message id: %s, payload: %s\n", o.name, m.id, m.payload)

}
