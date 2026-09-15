package queue

import "jobqueue/internal/job"

type Queue struct {
	jobs chan *job.Job
}

// Constructor for job queue
func New(size int) *Queue {
	return &Queue{
		jobs: make(chan *job.Job, size),
	}
}

// Will define some methods here for queue

// 1. Enqueue  it will add items to queue

func (q *Queue) Enqueue(j *job.Job) {
	q.jobs <- j
}

// 2. Now we need a reciever for job

func (q *Queue) Jobs() <-chan *job.Job {
	return q.jobs
}

func (q *Queue) Close() {
	close(q.jobs)
}
