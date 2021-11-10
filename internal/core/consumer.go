package core

type Consumer struct {
	jobChan    chan interface{}
	ingestChan chan interface{}
}

func NewConsumer(jobChan, ingestChan chan interface{}) *Consumer {
	return &Consumer{
		jobChan:    jobChan,
		ingestChan: ingestChan,
	}
}

func (c Consumer) Callback(event interface{}) {
	c.ingestChan <- event
}
