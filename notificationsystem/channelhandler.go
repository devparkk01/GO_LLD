package main

import "fmt"

type ChannelHandler interface {
	Send(*NotificationRequest) error
}

type EmailHandler struct{}

func (e *EmailHandler) Send(req *NotificationRequest) error {
	payload := req.GetPayload().(*EmailPayload)
	if payload == nil {
		return fmt.Errorf("does not conform to email payload")
	}
	fmt.Printf("[Email] To User %s | Subject: %s  | Body: %s\n", req.GetUserId(), payload.GetSubject(), payload.GetBody())
	return nil
}

type SMSHandler struct {}

func (e *SMSHandler) Send(req *NotificationRequest) error {
	payload := req.GetPayload().(*SMSPayload)
	if payload == nil {
		return fmt.Errorf("does not conform to SMS payload")
	}
	fmt.Printf("[SMS] To User %s | Text: %s\n", req.GetUserId(), payload.GetText())
	return nil
}

type PushHandler struct {}

func (e *PushHandler) Send(req *NotificationRequest) error {
	payload := req.GetPayload().(*PushPayload)
	if payload == nil {
		return fmt.Errorf("does not conform to Push payload")
	}
	fmt.Printf("[PUSH] To User %s | Title: %s  | Link: %s  | Body: %s\n", req.GetUserId(), payload.GetTitle(), payload.GetLink(), payload.GetBody())
	return nil
}
