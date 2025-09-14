package main

import (
	"fmt"
	"time"
)

type Worker struct {
	id      int
	service *NotificationService
}

func GetNewWorker(id int, service *NotificationService) *Worker {
	return &Worker{
		id:      id,
		service: service,
	}
}

func (w *Worker) Start() {
	// start polling messages
	fmt.Printf("Worker %d started...\n", w.id)

	for req := range w.service.queue {
		// get the retry count of requests 

		handler, ok := w.service.handlers[req.channelType]

		if !ok {
			fmt.Printf("No handler found for channel %s\n", req.channelType)
		} else {
			err := handler.Send(req)
			// when there is an error while sending notification
			if err != nil {
				fmt.Printf("Failed to send notification id %s. Error: %s\n", req.id, err.Error())
				req.retryCount++ 
				// send it to dlq if the retry count exceeds
				if req.retryCount >= w.service.maxRetry {
					w.service.SendToDlq(req)
				} else {
					time.Sleep(time.Millisecond * 200) // small backoff 
					w.service.queue <- req 
				}
			}
		}
	}
}