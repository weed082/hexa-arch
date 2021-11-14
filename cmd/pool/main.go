package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	numberOfWorkers = 10
)

func main() {
	// create a consumer
	consumer := Consumer{
		ingestChan: make(chan JobEvent, 1),
		jobsChan:   make(chan JobEvent, numberOfWorkers),
	}

	// simulate external lib sending jobs
	producer := Producer{CallbackFunc: consumer.RegisterEventCallback}
	go producer.StartSimulation()

	// context for cancellation
	ctx, cancelFunc := context.WithCancel(context.Background())

	// use a wait group to shutdown program only when all jobs are done
	wg := &sync.WaitGroup{}

	// start the consumer
	go consumer.Start(ctx)

	// add num of workers to wg
	wg.Add(numberOfWorkers)

	// start the worker pool
	for i := 0; i < numberOfWorkers; i++ {
		go consumer.Worker(wg, i)
	}

	// we want to create a termination channel so we can shut down gracefully
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, syscall.SIGINT, syscall.SIGTERM)

	<-terminationChan

	fmt.Println("--- SHUTDOWN SIGNAL RECEIVED ---")
	cancelFunc() // send shutdown signal through the context
	wg.Wait()    // wait here until all workers are done

	fmt.Println("All workers finished")
}

// Producer simulates an external library that invokes the registered
// callback when it has new data for us once per 100ms
type Producer struct {
	CallbackFunc func(event JobEvent)
}

func (p Producer) StartSimulation() {
	eventIndex := 0
	for {
		p.CallbackFunc(JobEvent{
			Id:   fmt.Sprintf("%d", eventIndex),
			Type: fmt.Sprintf("Simulation %v", eventIndex),
		})
		log.Printf("job send : %d", eventIndex)
		eventIndex++
		time.Sleep(time.Millisecond * 100)
	}
}

type JobEvent struct {
	Id   string
	Type string
}

type Consumer struct {
	ingestChan chan JobEvent
	jobsChan   chan JobEvent
}

func (c Consumer) Start(ctx context.Context) {
	for {
		select {
		case job := <-c.ingestChan:
			log.Printf("job received : %s", job.Id)
			c.jobsChan <- job
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan")
			close(c.jobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}

// Externally register an event
func (c Consumer) RegisterEventCallback(event JobEvent) {
	c.ingestChan <- event
}

// starts a single worker function
func (c Consumer) Worker(wg *sync.WaitGroup, workerIndex int) {
	fmt.Printf("Worker %d starting\n", workerIndex)
	// defer will be called once the jobsChan closes
	defer wg.Done()
	// blocks until an event is received or channel is closed
	for event := range c.jobsChan {
		// handle events
		fmt.Printf("Worker %d started job %v type %v\n", workerIndex, event.Id, event.Type)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished job %v type %v\n", workerIndex, event.Id, event.Type)
	}
	fmt.Printf("Worker %d interrupted\n", workerIndex)
}
