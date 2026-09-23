package handler

import (
	"context"
	"jobqueue/internal/job"
)

type VideoOperation interface {
	Execute(ctx context.Context, payload *job.VideoPayload) error
}
