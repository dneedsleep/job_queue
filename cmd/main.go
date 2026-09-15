package main

import (
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
	"jobqueue/internal/worker"
	"time"
)

func main() {

	q := queue.New(10)

	w1 := worker.CreateWorker(1, q)
	w2 := worker.CreateWorker(2, q)

	go w1.Start()
	go w2.Start()

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

	time.Sleep(10 * time.Second)

}
