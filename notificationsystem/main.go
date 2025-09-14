package main

import (
	"fmt"
	"time"
)


func main () {
	service := NewNotificationService(10)


	// Start 2 workers 
	worker1 := GetNewWorker(1, service)
	worker2 := GetNewWorker(2, service)
	go worker1.Start()
	go worker2.Start() 


	reqs := []*NotificationRequest{
		{
			id: "1111", 
			userId: "user1",
			channelType: EMAIL,
			payload: &EmailPayload{
				subject: "HOLA",
				body: "How are you",
			},

		},
		{
			id: "1134", 
			userId: "user1",
			channelType: SMS,
			payload: &SMSPayload{
				text: "You have been charged",
			},
		},
		{
			id: "1456", 
			userId: "user2",
			channelType: PUSH,
			payload: &PushPayload{
				title: "You have been charged",
				link: "https://www.google.com",
				body: "Click this link",
			},
		},
		{
			id: "2378", 
			userId: "user10",
			channelType: EMAIL,
			payload: &EmailPayload{
				subject: "Hiring",
				body: "Let's talk about business",
			},
		},
		{
			id: "3322", 
			userId: "user3",
			channelType: EMAIL,
			payload: &EmailPayload{
				subject: "Memories",
				body: "You have an email",
			},
		},
		{
			id: "8744", 
			userId: "user3",
			channelType: SMS,
			payload: &SMSPayload{
				text: "We are designing notification system",
			},
		},
	}

	for _, req := range reqs {
		err := service.SendNotification(req)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	time.Sleep(2 * time.Second)

}