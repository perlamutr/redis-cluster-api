package handler

import (
	"net/http"
)

//	GetKeys returns list of all keys in cluster using pattern parameter
func (h *Handlers) GetKeys(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("pattern")
	cmd := h.Keys(p)
	h.log("Using pattern: '%s'", p)
	result, err := cmd.Result()
	if err != nil {
		h.error(&w, err)
	}
	h.respond(&w, result)
}

