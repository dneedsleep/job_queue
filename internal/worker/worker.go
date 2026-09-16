package worker

import (
	"errors"
	"fmt"
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	ID int
	q  *queue.Queue
}

func CreateWorker(ID int, q *queue.Queue) *Worker {
	return &Worker{
		ID: ID,
		q:  q,
	}
}

func (w *Worker) Execute(j *job.Job) error {
	fmt.Printf("Worker %d processing job %s\n", w.ID, j.ID)
	if err := j.StatusUpdate(job.Processing); err != nil {
		return err
	}
	Time := time.Duration(rand.Intn(10)+1) * time.Second
	time.Sleep(Time)
	if Time >= 7*time.Second {
		if err := j.StatusUpdate(job.Failed); err != nil {
			return err
		}
		return errors.New("Time taking to long")
	}

	if err := j.StatusUpdate(job.Completed); err != nil {
		return err
	}
	return nil
}

func (w *Worker) Start(wg, jobwg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d started\n", w.ID)

	for j := range w.q.Jobs() {
		err := w.Execute(j)
		if err != nil {
			fmt.Printf("Worker %d failed to complete     job %s\n", w.ID, j.ID)
			j.RetryCount++
			if j.RetryCount <= j.MaxRetries {
				fmt.Printf(
					"Job %s failed. Retry count: %d\n",
					j.ID,
					j.RetryCount,
				)
				j.StatusUpdate(job.Pending)
				w.q.Enqueue(j)
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
