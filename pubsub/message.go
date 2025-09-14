package main 

type Message struct {
	id string 
	payload string 
}

func NewMessage(id, payload string) *Message {
	return &Message{
		id: id, 
		payload: payload,
	}
}