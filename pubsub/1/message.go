package main

import "time"

type Message struct {
	Id        string
	Payload   string
	Topic     string
	Timestamp time.Time
}

func NewMessage(id, payload, topic string, timestamp time.Time) *Message {
	return &Message{
		Id:        id,
		Payload:   payload,
		Topic:     topic,
		Timestamp: timestamp,
	}
}
