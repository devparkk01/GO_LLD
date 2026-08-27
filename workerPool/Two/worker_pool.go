package main

import (
	"fmt"
	"sync"
)

type JobFunc func(id int) string

type Job struct {
	id      int
	jobFunc JobFunc
}

type WorkerPool struct {
	concurrency int
	jobsChan    chan Job
	wg          sync.WaitGroup
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		concurrency: concurrency,
		jobsChan:    make(chan Job, 3),
	}
}

func (wp *WorkerPool) Submit(jobFunc JobFunc, id int) {
	wp.jobsChan <- Job{id: id, jobFunc: jobFunc}
}

func (wp *WorkerPool) Run() {
	// wg := &sync.WaitGroup{}
	for range wp.concurrency {
		// fire worker to do the job
		wp.wg.Add(1)
		go wp.worker()
	}
	wp.wg.Wait()
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	// fmt.Printf("%v\n", wg)
	for job := range wp.jobsChan {
		fmt.Println(job.jobFunc(job.id))
		// fmt.Println(job())
	}
}

func (wp *WorkerPool) Done() {
	close(wp.jobsChan)
}
