package worker

import (
	"context"
	"jobqueue/internal/job"
)

type Handler interface {
	Handle(ctx context.Context, j *job.Job) error
}
