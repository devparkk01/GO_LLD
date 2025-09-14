package main

import (
	"fmt"
	// "os"
	// "encoding/json"
)

type NotificationService struct {
	handlers map[ChannelType]ChannelHandler
	queue    chan *NotificationRequest
	dlq chan *NotificationRequest
	maxRetry int 
}

func NewNotificationService(queueSize int) *NotificationService {
	return &NotificationService{
		handlers: map[ChannelType]ChannelHandler{
			EMAIL: &EmailHandler{},
			SMS:   &SMSHandler{},
			PUSH:  &PushHandler{},
		},
		queue: make(chan *NotificationRequest, queueSize),
		dlq: make(chan *NotificationRequest, queueSize),
		maxRetry: 3,
	}
}

func (n *NotificationService) SendNotification(req *NotificationRequest) error {
	err := req.Validate()
	if err != nil {
		return err
	}
	fmt.Printf("Publishing notification id %s to the queue.\n", req.id)
	// add notification to the queue
	n.queue <- req 
	return nil
}

func (n *NotificationService) SendToDlq(req *NotificationRequest) {
	fmt.Println("Moved to DLQ: ", req.GetId())
	n.dlq <- req 

	// if we want to persist it to file 
	// f, _ := os.OpenFile("dlq.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	// defer f.Close() 
	// data, _ := json.Marshal(req)
	// f.WriteString(string(data) + "\n")
}