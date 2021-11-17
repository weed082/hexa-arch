package concurrency

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type WorkerPool struct {
	wg       *sync.WaitGroup
	ctx      context.Context
	jobCh    chan Job
	ingestCh chan Job
}

func NewWorkerPool(wg *sync.WaitGroup, ctx context.Context, jobChan chan Job, ingestChan chan Job) *WorkerPool {
	return &WorkerPool{
		wg:       wg,
		ctx:      ctx,
		jobCh:    jobChan,
		ingestCh: ingestChan,
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

//! (2) start worker pool
func (c *WorkerPool) Start() {
	for {
		select {
		case job := <-c.ingestCh:
			c.jobCh <- job
		case <-c.ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobChan")
			close(c.jobCh)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}

//! (3) register event callback
func (c *WorkerPool) RegisterJobCallback(job Job) {
	c.ingestCh <- job
}
