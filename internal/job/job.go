package job

import (
	"errors"
	"sync"
)

type Status string

const (
	Pending    Status = "pending"
	Processing Status = "processing"
	Completed  Status = "completed"
	Failed     Status = "failed"
)

type Job struct {
	ID         string
	Type       string
	Payload    any
	Status     Status
	RetryCount int
	MaxRetries int
	mu         sync.RWMutex
}

func (j *Job) StatusUpdate(UpdatedStatus Status) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	if UpdatedStatus == Processing && j.Status != Pending {
		return errors.New("Status should be pending")
	} else if (UpdatedStatus == Completed || UpdatedStatus == Failed) && j.Status != Processing {
		return errors.New("Status should be processing")
	} else if UpdatedStatus == Pending && j.Status != Failed {
		return errors.New("Status cannot be changed to pending")
	}
	j.Status = UpdatedStatus
	return nil

}
