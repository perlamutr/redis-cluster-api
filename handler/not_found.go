package handler

import (
	"net/http"
)

const INSTRUCTION = "Please use following URI with methods: " +
	"GET /keys?pattern=<PATTERN> to get list of keys in cluster ; " +
	"POST /set/{key} to set any key (use method body for value) ; " +
	"POST /lpush/{queue} to add json to queue (json in method body ; " +
	"GET /lrange/{queue} to get list of queue"

type NotFound struct {
	*Handlers
}

func NewNotFound(h *Handlers) *NotFound {
	return &NotFound{h}
}

//	ServeHTTP is an action for MethodNotFound error
func (h *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log("Method not found. Request: %#v", r)
	h.error(&w, "Sorry such method is absent. Try to read logs for details. " + INSTRUCTION)
}


