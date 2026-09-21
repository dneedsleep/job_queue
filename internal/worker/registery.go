package worker

type HandlerRegistry struct {
	handlers map[string]Handler
}

func NewRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		handlers: make(map[string]Handler),
	}
}

func (r *HandlerRegistry) Register(jobType string, handler Handler) {
	r.handlers[jobType] = handler
}

func (r *HandlerRegistry) Get(jobType string) (Handler, bool) {

	handler, ok := r.handlers[jobType]
	return handler, ok
}
