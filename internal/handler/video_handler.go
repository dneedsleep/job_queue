package handler

import (
	"context"
	"errors"
	"fmt"
	"jobqueue/internal/job"
	"os/exec"
)

type VideoHandler struct {
}

func (v *VideoHandler) Handle(ctx context.Context, j *job.Job) error {

	payload, ok := j.Payload.(job.VideoPayload)
	if !ok {
		return errors.New("invalid video payload")
	}

	switch payload.Operation {
	case "resize":
		return v.Resize(ctx, &payload)
	default:
		return fmt.Errorf("Operation not implemented")
	}

}

func (v *VideoHandler) Resize(ctx context.Context, payload *job.VideoPayload) error {

	cmd := exec.CommandContext(
		ctx,
		"ffmpeg",
		"-i",
		payload.InputPath,
		"-vf", fmt.Sprintf(
			"scale=%d:%d",
			payload.Width,
			payload.Height,
		),
		payload.OutputPath,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"ffmpeg resize failed: %w\n%s",
			err,
			output,
		)
	}

	fmt.Printf(
		"Processing video: %s -> %s\n",
		payload.InputPath,
		payload.OutputPath,
	)

	return nil
}
