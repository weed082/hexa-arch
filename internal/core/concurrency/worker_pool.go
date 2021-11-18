package concurrency

import (
	"context"
	"log"
	"reflect"
	"sync"
)

type WorkerPool struct {
	wg          *sync.WaitGroup
	ctx         context.Context
	jobChs      []chan Job
	ingestChs   []chan Job
	ingestCases []reflect.SelectCase
}

func NewWorkerPool(wg *sync.WaitGroup, ctx context.Context) *WorkerPool {
	return &WorkerPool{
		wg:          wg,
		ctx:         ctx,
		jobChs:      []chan Job{},
		ingestChs:   []chan Job{},
		ingestCases: []reflect.SelectCase{},
	}
}

func (wp *WorkerPool) Start() {
	for {
		select {
		case <-wp.ctx.Done():
			log.Println("close signal received, closing job channels")
			for _, ch := range wp.jobChs {
				close(ch)
			}
			log.Println("closed all job channels")
			return
		default:
			log.Println(wp.ingestCases)
			chosenIdx, job, ok := reflect.Select(wp.ingestCases)
			if !ok {
				log.Println("received at closing chan")
				continue
			}
			log.Println("job send")
			wp.jobChs[chosenIdx] <- job.Interface().(Job)
		}
	}
}

func (wp *WorkerPool) AddPool(workerCnt int) int {
	newPoolIdx := len(wp.jobChs) // new pool index
	ingestCh := make(chan Job)

	// add job, ingest channels
	wp.ingestCases = append(wp.ingestCases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ingestCh)})
	wp.ingestChs = append(wp.ingestChs, ingestCh)
	wp.jobChs = append(wp.jobChs, make(chan Job))
	log.Println(wp.ingestChs)

	// start workers
	wp.wg.Add(workerCnt)
	for i := 0; i < workerCnt; i++ {
		go wp.work(newPoolIdx)
	}
	return newPoolIdx
}

func (wp *WorkerPool) work(poolIdx int) {
	defer wp.wg.Done()
	log.Println(wp.jobChs[poolIdx])
	for job := range wp.jobChs[poolIdx] {
		log.Println("working")
		job.Callback()
	}
	log.Println("worker destroyed")
}

func (wp *WorkerPool) RegisterJob(poolIdx int, job Job) {
	log.Println("registered")
	log.Println(wp.ingestChs[poolIdx])
	wp.ingestChs[poolIdx] <- job
}
