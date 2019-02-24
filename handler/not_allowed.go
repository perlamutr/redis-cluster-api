package handler

import (
	"net/http"
)

type NotAllowed struct {
	*Handlers
}

func NewNotAllowed(h *Handlers) *NotAllowed {
	return &NotAllowed{h}
}

//	ServeHTTP is an action for MethodNotAllowed error
func (h *NotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log("Method not allowed. Request: %#v", r)
	h.error(&w, "Sorry such method not allowed. Try to read logs for details")
}
