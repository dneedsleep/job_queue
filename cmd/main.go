package main

import (
	"context"
	"jobqueue/internal/handler"
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
	"jobqueue/internal/worker"
	"sync"
)

func main() {

	q := queue.New(100)

	registry := worker.NewRegistry()

	videoHandler := &handler.VideoHandler{}

	registry.Register("video", videoHandler)

	var wg sync.WaitGroup
	var jobWg sync.WaitGroup

	jobWg.Add(4)
	wg.Add(2)

	ctx := context.Background()

	w1 := worker.CreateWorker(1, q, registry)
	w2 := worker.CreateWorker(2, q, registry)

	go w1.Start(ctx, &wg, &jobWg)
	go w2.Start(ctx, &wg, &jobWg)

	q.Enqueue(&job.Job{
		ID:         "1",
		Type:       "video",
		Status:     job.Pending,
		MaxRetries: 2,
	})

	q.Enqueue(&job.Job{
		ID:         "2",
		Type:       "video",
		Status:     job.Pending,
		MaxRetries: 2,
	})

	q.Enqueue(&job.Job{
		ID:         "3",
		Type:       "video",
		Status:     job.Pending,
		MaxRetries: 2,
	})

	q.Enqueue(&job.Job{
		ID:         "4",
		Type:       "video",
		Status:     job.Pending,
		MaxRetries: 2,
	})
	//q.Close()
	jobWg.Wait()
	q.Close()

	wg.Wait()
}
