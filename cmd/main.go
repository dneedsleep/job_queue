package main

import (
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
	"jobqueue/internal/worker"
	"sync"
)

func main() {

	q := queue.New(100)

	var wg sync.WaitGroup
	var jobWg sync.WaitGroup

	jobWg.Add(4)
	wg.Add(2)

	w1 := worker.CreateWorker(1, q)
	w2 := worker.CreateWorker(2, q)

	go w1.Start(&wg, &jobWg)
	go w2.Start(&wg, &jobWg)

	q.Enqueue(&job.Job{
		ID:         "1",
		Type:       "test",
		Status:     job.Pending,
		MaxRetries: 2,
	})

	q.Enqueue(&job.Job{
		ID:         "2",
		Type:       "test",
		Status:     job.Pending,
		MaxRetries: 2,
	})

	q.Enqueue(&job.Job{
		ID:         "3",
		Type:       "test",
		Status:     job.Pending,
		MaxRetries: 2,
	})

	q.Enqueue(&job.Job{
		ID:         "4",
		Type:       "test",
		Status:     job.Pending,
		MaxRetries: 2,
	})
	//q.Close()
	jobWg.Wait()
	q.Close()

	wg.Wait()
}
