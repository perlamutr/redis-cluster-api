package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//	LPush adds json record to selected queue
func (h *Handlers) LPush(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	queue := p["queue"]
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.error(&w, err)
		return
	}
	h.log("value:%#v", value)
	var v interface{}
	err = json.Unmarshal(value, &v)
	if err != nil {
		h.error(&w, err)
		return
	}

	cmd := h.ClusterClient.LPush(queue, value)
	h.log("Adding to Queue: '%s' Value: '%s'", queue, value)
	result, err := cmd.Result()
	if err != nil {
		h.error(&w, err)
	}
	h.respond(&w, map[string]int64{"in_queue": result})
}
