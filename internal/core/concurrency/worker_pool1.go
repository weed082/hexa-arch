package concurrency

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type WorkerPool1 struct {
	wg       *sync.WaitGroup
	ctx      context.Context
	jobCh    chan Job
	ingestCh chan Job
}

func NewWoo(wg *sync.WaitGroup, ctx context.Context, jobChan chan Job, ingestChan chan Job) *WorkerPool1 {
	return &WorkerPool1{
		wg:       wg,
		ctx:      ctx,
		jobCh:    jobChan,
		ingestCh: ingestChan,
	}
}

//! (1) generate worker
func (c *WorkerPool1) Generate(count int) {
	c.wg.Add(count)
	for i := 0; i < count; i++ {
		go c.work()
	}
}

func (c *WorkerPool1) work() {
	defer c.wg.Done()
	for job := range c.jobCh {
		job.Callback()
	}
	log.Println("worker destroyed")
}

//! (2) start worker pool
func (c *WorkerPool1) Start() {
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
func (c *WorkerPool1) RegisterJobCallback(job Job) {
	c.ingestCh <- job
}
