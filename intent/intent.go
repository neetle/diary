package intent

import "fmt"

type Intent interface{}

type None struct{}
type ContentUpdated string
type ContentWrite string
type Quit struct{}

type FailedWrite error

func NewHandler() *Handler {
	write := make(chan string, 1)
	quit := make(chan struct{})

	return &Handler{
		contentWrite: write,
		ContentWrite: write,
		quit:         quit,
		Quit:         quit,
	}
}

type Handler struct {
	contentWrite chan<- string
	quit         chan<- struct{}

	ContentWrite <-chan string
	Quit         <-chan struct{}
}

func (h *Handler) Publish(intents ...Intent) {
	for i := 0; i < len(intents); i++ {
		h.delegate(intents[i])
	}
}

func (h *Handler) delegate(process Intent) {
	select {
	case <-h.Quit:
		return
	default:
	}

	switch p := process.(type) {
	case Quit:
		close(h.quit)
		return

	case ContentUpdated:
		if len(p) >= 2 && p[len(p)-2:] == "\n\n" {
			h.contentWrite <- string(p[:len(p)-2])
			close(h.contentWrite)
		}
		return
	}

	panic(fmt.Sprintf("UNKNOWN INTENT %T", process))
}
