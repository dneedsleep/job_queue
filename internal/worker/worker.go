package worker

import (
	"context"
	"fmt"
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
	"sync"
	"time"
)

type Worker struct {
	ID       int
	q        *queue.Queue
	registry *HandlerRegistry
}

func CreateWorker(ID int, q *queue.Queue, registry *HandlerRegistry) *Worker {
	return &Worker{
		ID:       ID,
		q:        q,
		registry: registry,
	}
}

func (w *Worker) Start(ctx context.Context, wg, jobwg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d started\n", w.ID)

	for j := range w.q.Jobs() {

		handler, ok := w.registry.Get(j.Type)

		if !ok {
			fmt.Printf("No handler found for job type: %s\n", j.Type)
			j.StatusUpdate(job.Failed)
			jobwg.Done()
			continue
		}

		if err := j.StatusUpdate(job.Processing); err != nil {
			fmt.Printf("Failed to update job %s to processing: %v\n", j.ID, err)
			jobwg.Done()
			continue
		}

		start := time.Now()
		err := handler.Handle(ctx, j)

		duration := time.Since(start)

		// old logic
		//err := w.Execute(ctx, j)
		if err != nil {
			fmt.Printf("Worker %d failed to complete     job %s\n", w.ID, j.ID)
			j.RetryCount++

			if j.RetryCount <= j.MaxRetries {

				delay := time.Duration(j.RetryCount) * time.Second

				fmt.Printf(
					"Job %s failed. Retry count: %d\n",
					j.ID,
					j.RetryCount,
				)

				j.StatusUpdate(job.Pending)

				// time AFERER func function fires a go rountine that it will be implement after this much time
				time.AfterFunc(delay, func() {
					w.q.Enqueue(j)
				})
				continue
			}

			fmt.Printf(
				"Job %s permanently failed after %d retries\n",
				j.ID,
				j.RetryCount,
			)
			jobwg.Done()
			continue

		}
		fmt.Printf("Worker %d completed job , time taken %v %s\n", w.ID, duration, j.ID)
		jobwg.Done()

	}
}
