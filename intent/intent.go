package intent

import (
	"fyne.io/fyne"
)

type Intent interface{}

type None struct{}
type Start struct{}
type ContentUpdated string
type ContentWrite string
type Quit struct{}

type FailedWrite error

type Handler struct {
	App      fyne.App
	Delegate func(fyne.App, *Handler, Intent) Intent
}

func (h *Handler) Publish(intents ...Intent) {
	for i := 0; i < len(intents); i++ {
		next := h.Delegate(h.App, h, intents[i])
		if _, ok := next.(None); !ok {
			intents = append(intents, next)
		}
	}
}
