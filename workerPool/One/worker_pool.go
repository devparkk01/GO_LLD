package main

import (
	"sync"
)

type WorkerPool struct {
	concurrency int
	wg          sync.WaitGroup
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		concurrency: concurrency,
	}
}

func (w *WorkerPool) Run(jobsChan <-chan Job, resultsChan chan<- Result) {
	for range w.concurrency {
		w.wg.Add(1)
		go w.worker(jobsChan, resultsChan)
	}
	// wait for all workers to finish execution
	w.wg.Wait()
	// close resultsChan. 
	// workerpool is producer for resultsChan. so it must be closed by workerPool
	close(resultsChan)
}

func (w *WorkerPool) worker(jobsChan <-chan Job, resultsChan chan<- Result) {
	defer w.wg.Done()
	for job := range jobsChan {
		out, err := job.Process()
		resultsChan <- Result{out, err}
	}
}
