package main

import "fmt"

type ChannelType string

const (
	SMS   ChannelType = "sms"
	EMAIL ChannelType = "email"
	PUSH  ChannelType = "push"
)

type NotificationPayload interface {
	GetMessage() string
}
// Email Payload 
type EmailPayload struct {
	subject string
	body    string
	cc      []string
}

func (e *EmailPayload) GetMessage() string {
	return fmt.Sprintf("%s:%s", e.subject, e.body)
}

func (e *EmailPayload) GetSubject() string {
	return e.subject
}

func (e *EmailPayload) GetBody() string {
	return e.body
}

func (e *EmailPayload) GetCC() []string {
	return e.cc 
}

// SMS payload 
type SMSPayload struct {
	text string
}

func (s *SMSPayload) GetMessage() string {
	return s.text 
}

func (s *SMSPayload) GetText() string {
	return s.text
}

// Push payload 
type PushPayload struct {
	title string
	link  string
	body  string
}

func (p *PushPayload) GetMessage() string {
	return fmt.Sprintf("%s:%s:%s", p.title, p.link, p.body)
}
func (p *PushPayload) GetTitle() string {
	return p.title 
}
func (p *PushPayload) GetLink() string {
	return p.link 
}
func (p *PushPayload) GetBody() string {
	return p.body 
}


// ----------------------------------
// NotificationRequest
// ----------------------------------

type NotificationRequest struct {
	id          string
	userId      string
	channelType ChannelType
	payload     NotificationPayload
	retryCount int 
}

func (n *NotificationRequest) GetId() string {
	return n.id
}
func (n *NotificationRequest) GetUserId() string {
	return n.userId
}
func (n *NotificationRequest) GetChannelType() ChannelType {
	return n.channelType
}
func (n *NotificationRequest) GetPayload() NotificationPayload {
	return n.payload
}
func (n *NotificationRequest) GetRetryCount() int {
	return n.retryCount
}

func (n *NotificationRequest) Validate() error {
	switch n.channelType {
	case EMAIL:
		_, ok := n.payload.(*EmailPayload)
		if !ok {
			return fmt.Errorf("invalid payload for Email channel")
		}
	case SMS:
		_, ok := n.payload.(*SMSPayload) 
		if !ok {
			return fmt.Errorf("invalid payload for SMS channel")
		}
	case PUSH:
		_, ok := n.payload.(*PushPayload) 
		if !ok {
			return fmt.Errorf("invalid payload for Push channal")
		}
	default:
		return fmt.Errorf("unknown channel type %s", n.channelType)
	}
	return nil 
	
}