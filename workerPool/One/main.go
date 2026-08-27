package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	workerPool := NewWorkerPool(2)
	jobs := []Job{
		&Email{Id: 1, ReceiverMail: "a@gmail.com"},
		&Email{Id: 2, ReceiverMail: "b"},
		&Message{Id: 3, Receipient: "alpha"},
		&Email{Id: 4, ReceiverMail: "c@gmail.com"},
		&Message{Id: 5, Receipient: "beta"},
		&Email{Id: 6, ReceiverMail: "d@gmail.com"},
		&Message{Id: 7, Receipient: "gama"},
		&Email{Id: 8, ReceiverMail: "e@gmail.com"},
		&Message{Id: 9, Receipient: "delta"}, 
		&Email{Id: 10, ReceiverMail: "f@gmail.com"},
	}

	jobsChan := make(chan Job, 3)
	resultsChan := make(chan Result)
	wg := sync.WaitGroup{}


	//  start putting items into jobsChan 
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, job := range jobs {
			jobsChan <- job 
		}
		// close the channel when it's done producing 
		// Main function is producer for JobsChan, so it must be closed by main function
		close(jobsChan) // signals workers that there are no more jobs 
	}()

	// start consuming from resultsChan 
	wg.Add(1)
	go func() {
		defer wg.Done() 
		for result := range resultsChan {
			time.Sleep(1 * time.Second)
			if result.Err != nil {
				fmt.Printf("Error: %v\n" , result.Err)
			} else {
				fmt.Printf("Result: %v\n", result.Output)
			}
		}
	}()
	// start all workers 
	wg.Add(1)
	go func() {
		defer wg.Done()
		workerPool.Run(jobsChan, resultsChan)
	}()
	// wait for channels to finish consuming 
	wg.Wait()
}
