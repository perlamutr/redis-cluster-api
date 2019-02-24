package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"redis-responder/storage"
	"time"
)

type Handlers struct {
	*storage.Redis
}

func NewHandlers(redis *storage.Redis) *Handlers {
	return &Handlers{redis}
}

//	AddRoutes setting application routes
func (h *Handlers) AddRoutes(router *mux.Router) {
	router.HandleFunc("/keys", h.GetKeys).Methods("GET")
	router.HandleFunc("/set/{key}", h.Set).Methods("POST")
	router.HandleFunc("/lpush/{queue}", h.LPush).Methods("POST")
	router.HandleFunc("/lrange/{queue}", h.LRange).Methods("GET")
	router.MethodNotAllowedHandler = NewNotAllowed(h)
	router.NotFoundHandler = NewNotFound(h)
}

func (*Handlers) error(w *http.ResponseWriter, msg interface{}) {
	json.NewEncoder(*w).Encode(map[string]string{"error": fmt.Sprint(msg)})
}

func (*Handlers) respond(w *http.ResponseWriter, msg interface{}) {
	bytes, _ := json.Marshal(map[string]interface{}{"result": msg})
	(*w).Write(bytes)
}

func (*Handlers) log(format string, msg ...interface{})  {
	m := make([]interface{}, 0, len(msg) + 1)
	m = append(m, time.Now().Format("2006-01-02"))
	m = append(m, msg...)
	fmt.Printf("[%s] " + format + "\n", m...)
}
