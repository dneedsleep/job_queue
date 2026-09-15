package job

type Status string

const (
	Pending    Status = "pending"
	Processing Status = "processing"
	Completed  Status = "completed"
	Failed     Status = "failed"
)

type Job struct {
	ID      string
	Type    string
	Payload string
	Status  Status
}
