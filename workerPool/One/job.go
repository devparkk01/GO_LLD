package main

import "fmt"

type Job interface {
	Process() (string, error)
}

type Email struct {
	Id           int
	ReceiverMail string
}

func (e *Email) Process() (string, error) {
	out := fmt.Sprintf("Processing %d, Sending email to %s", e.Id, e.ReceiverMail)
	return out, nil
}

type Message struct {
	Id         int
	Receipient string
}

func (m *Message) Process() (string, error) {
	out := fmt.Sprintf("Processing %d, Sending message to %s", m.Id, m.Receipient)
	return out, nil
}

type Result struct {
	Output string
	Err    error
}
