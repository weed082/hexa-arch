package pool

import (
	"log"
	"sync"
)

type WorkerPool struct {
	logger *log.Logger
	wg     *sync.WaitGroup
	jobCh  chan func()
}

func NewWorkerPool(logger *log.Logger, wg *sync.WaitGroup, bufferSize int) *WorkerPool {
	return &WorkerPool{
		logger: logger,
		wg:     &sync.WaitGroup{},
		jobCh:  make(chan func(), bufferSize),
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
	for callback := range wp.jobCh {
		callback()
	}
	wp.logger.Println("worker destroyed")
}

func (wp *WorkerPool) RegisterJob(callback func()) {
	wp.jobCh <- callback
}

func (wp *WorkerPool) Stop() {
	close(wp.jobCh)
}
