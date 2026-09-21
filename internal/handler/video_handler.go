package handler

import (
	"context"
	"errors"
	"fmt"
	"jobqueue/internal/job"
)

type VideoHandler struct {
}

func (v *VideoHandler) Handle(ctx context.Context, j *job.Job) error {

	payload, ok := j.Payload.(job.VideoPayload)
	if !ok {
		return errors.New("invalid video payload")
	}

	fmt.Printf(
		"Processing video: %s -> %s\n",
		payload.InputPath,
		payload.OutputPath,
	)

	return nil
}
