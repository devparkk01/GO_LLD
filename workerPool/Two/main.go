package main

import (
	"fmt"
	"sync"
)

func main() {
	wp := NewWorkerPool(2)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		wp.Run()
	}()

	wg.Add(1)
	go func() {
		// signals workerpool that client(main) is done producing items to the workerpool
		defer wp.Done()
		defer wg.Done()
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed Email ", id)
		}, 1)
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed Message", id)
		}, 2)
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed Email", id)
		}, 3)
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed JOB 4", id)
		}, 4)
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed JOB 5", id)
		}, 5)
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed JOB 6", id)
		}, 6)
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed JOB 7", id)
		}, 7)
		wp.Submit(func(id int) string {
			return fmt.Sprintf("id: %d, Processed JOB 8", id)
		}, 8)
	}()

	wg.Wait()
}
