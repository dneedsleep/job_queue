package job

type VideoPayload struct {
	Operation  string
	InputPath  string
	OutputPath string
	Height     int
	Width      int
}
