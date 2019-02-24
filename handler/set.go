package handler

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//	Set value to selected key with expiration value (infinite by default)
func (h *Handlers) Set(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["key"]
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.error(&w, err)
		return
	}
	var exp time.Duration
	if val := r.Form.Get("expiration"); val != "" {
		i, _ := strconv.ParseInt(val, 10, 32)
		exp = time.Duration(i) * time.Second
	}
	h.log("Setting key %s to '%s'", key, value)
	cmd := h.Redis.Set(key, value, exp)
	result, err := cmd.Result()
	if err != nil {
		h.error(&w, err)
	}
	h.respond(&w, result)
}
