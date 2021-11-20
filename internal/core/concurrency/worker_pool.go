package concurrency

import (
	"log"
	"sync"
)

type WorkerPool struct {
	wg    *sync.WaitGroup
	jobCh chan Job
}

func NewWorkerPool(wg *sync.WaitGroup, buffuerSize int) *WorkerPool {
	return &WorkerPool{
		wg:    wg,
		jobCh: make(chan Job, buffuerSize),
	}
}

//! (2) generate worker
func (wp *WorkerPool) Generate(count int) {
	wp.wg.Add(count)
	for i := 0; i < count; i++ {
		go wp.work()
	}
}

func (wp *WorkerPool) work() {
	defer wp.wg.Done()
	for job := range wp.jobCh {
		job.Callback()
	}
	log.Println("worker destroyed")
}

func (wp *WorkerPool) RegisterJob(callback func()) {
	wp.jobCh <- Job{callback}
}

func (wp *WorkerPool) Close() {
	close(wp.jobCh)
}
