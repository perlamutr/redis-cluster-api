package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

//	LRang checks records in selected queue
func (h *Handlers) LRange(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	queue := p["queue"]
	cmd := h.ClusterClient.LRange(queue, 0, -1)
	h.log("Reading queue Queue: '%s'", queue)
	result, err := cmd.Result()
	if err != nil {
		h.error(&w, err)
	}
	h.respond(&w, result)
}
