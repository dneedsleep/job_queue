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

	jobWg.Add(1)
	wg.Add(1)

	ctx := context.Background()

	w1 := worker.CreateWorker(1, q, registry)
	//w2 := worker.CreateWorker(2, q, registry)

	go w1.Start(ctx, &wg, &jobWg)
	//go w2.Start(ctx, &wg, &jobWg)

	videoPayload := job.VideoPayload{
		Operation:  "resize",
		InputPath:  "videos/input.mp4",
		OutputPath: "videos/output.mp4",
		Width:      1280,
		Height:     720,
	}

	videoJob := &job.Job{
		ID:         "1",
		Type:       "video",
		Payload:    videoPayload,
		Status:     job.Pending,
		MaxRetries: 3,
	}

	q.Enqueue(videoJob)
	//q.Close()
	jobWg.Wait()
	q.Close()

	wg.Wait()
}
