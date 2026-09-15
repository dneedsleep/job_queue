package main

import (
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
	"jobqueue/internal/worker"
	"sync"
)

func main() {

	q := queue.New(10)

	var wg sync.WaitGroup

	wg.Add(2)

	w1 := worker.CreateWorker(1, q)
	w2 := worker.CreateWorker(2, q)

	go w1.Start(&wg)
	go w2.Start(&wg)

	q.Enqueue(job.Job{
		ID:   "1",
		Type: "test",
	})

	q.Enqueue(job.Job{
		ID:   "2",
		Type: "test",
	})

	q.Enqueue(job.Job{
		ID:   "3",
		Type: "test",
	})

	q.Enqueue(job.Job{
		ID:   "4",
		Type: "test",
	})
	q.Close()
	wg.Wait()
}
