package concurrency

import (
	"fmt"
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

//! (1) generate worker
func (c *WorkerPool) Generate(count int) {
	c.wg.Add(count)
	for i := 0; i < count; i++ {
		go c.work()
	}
}

func (c *WorkerPool) work() {
	defer c.wg.Done()
	for job := range c.jobCh {
		job.Callback()
	}
	log.Println("worker destroyed")
}

func (c *WorkerPool) RegisterJob(callback func()) {
	c.jobCh <- Job{callback}
}

func (c *WorkerPool) GracefulShutdown() {
	close(c.jobCh)
	fmt.Println("worker pool closing jobs channel")
}
