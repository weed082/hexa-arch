package concurrency

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type WorkerPool struct {
	wg         *sync.WaitGroup
	ctx        context.Context
	jobChan    chan Job
	ingestChan chan Job
}

func NewWorkerPool(wg *sync.WaitGroup, ctx context.Context, jobChan chan Job, ingestChan chan Job) *WorkerPool {
	return &WorkerPool{
		wg:         wg,
		ctx:        ctx,
		jobChan:    jobChan,
		ingestChan: ingestChan,
	}
}

//! (1) generate worker
func (c WorkerPool) Generate(count int) {
	c.wg.Add(count)
	for i := 0; i < count; i++ {
		go c.work()
	}
}

func (c WorkerPool) work() {
	defer c.wg.Done()
	for job := range c.jobChan {
		job.Callback()
	}
	log.Println("worker destroyed")
}

//! (2) start worker pool
func (c WorkerPool) Start() {
	for {
		select {
		case job := <-c.ingestChan:
			log.Println("job ingested")
			c.jobChan <- job
		case <-c.ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobChan")
			close(c.jobChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}

//! (3) register event callback
func (c WorkerPool) RegisterJobCallback(job Job) {
	c.ingestChan <- job
}
