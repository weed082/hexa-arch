package concurrency

import "sync"

type Job struct {
}

type Consumer struct {
	jobChan    chan Job
	ingestChan chan interface{}
}

func NewConsumer(jobChan chan Job, ingestChan chan interface{}) *Consumer {
	return &Consumer{
		jobChan:    jobChan,
		ingestChan: ingestChan,
	}
}

func (c Consumer) Generate(wg sync.WaitGroup, count int) {
	for i := 0; i < 3; i++ {
		go c.work()
	}
}

func (c Consumer) work() {

}
