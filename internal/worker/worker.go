package worker

import (
	"context"
	"fmt"
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
	"math/rand"
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

func (w *Worker) Execute(prtctx context.Context, j *job.Job) error {

	ctx, cancel := context.WithTimeout(prtctx, 5*time.Second)

	defer cancel()

	fmt.Printf("Worker %d processing job %s\n", w.ID, j.ID)
	if err := j.StatusUpdate(job.Processing); err != nil {
		return err
	}
	executionTime := time.Duration(rand.Intn(10)+1) * time.Second
	select {
	case <-time.After(executionTime):
		// simulated work finished

	case <-ctx.Done():
		return ctx.Err()
	}

	if err := j.StatusUpdate(job.Completed); err != nil {
		return err
	}
	return nil
}

func (w *Worker) Start(ctx context.Context, wg, jobwg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d started\n", w.ID)

	for j := range w.q.Jobs() {

		handler, ok := w.registry.Get(j.Type)

		if !ok {
			fmt.Printf("No handler found for job type: %s\n", j.Type)
			// handle failure/retry here later
			continue
		}

		err := handler.Handle(ctx, j)

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
		fmt.Printf("Worker %d completed     job %s\n", w.ID, j.ID)
		jobwg.Done()

	}
}
