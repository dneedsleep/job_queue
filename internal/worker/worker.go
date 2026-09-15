package worker

import (
	"fmt"
	"jobqueue/internal/job"
	"jobqueue/internal/queue"
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

func (w *Worker) Process(j job.Job) {
	fmt.Printf("Worker %d processing job %s\n", w.ID, j.ID)
	time.Sleep(5 * time.Second)
	fmt.Printf("Worker %d processed job %s\n", w.ID, j.ID)
}

func (w *Worker) Start() {
	fmt.Printf("Worker %d started\n", w.ID)

	for j := range w.q.Jobs() {
		fmt.Printf("Worker %d picked Job %s\n", w.ID, j.ID)

		time.Sleep(2 * time.Second)

		fmt.Printf("Worker %d completed Job %s \n", w.ID, j.ID)
	}
}
