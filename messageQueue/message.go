package main

import "time"

type Message struct {
	ID        string
	Payload   []byte
	CreatedAt time.Time
}

func NewMessage(id string, payload []byte) Message {
	return Message{
		ID:        id,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}
