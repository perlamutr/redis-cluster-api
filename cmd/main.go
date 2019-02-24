package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"redis-responder/handler"
	"redis-responder/storage"
	"strconv"
	"time"
)

func main() {
	host := os.Getenv("REDIS_HOST")
	reconnect := os.Getenv("RECONNECT_SEC")
	reconnectSec, err := strconv.ParseInt(reconnect, 10, 32)
	if err != nil {
		reconnectSec = 0
	}
	redis := storage.NewRedis()
	for {
		err := redis.Connect(host)
		if err == nil {
			break
		}
		log.Printf("Cannot connect redis cluster (%s): %s\n", host, err)
		if reconnectSec == 0 {
			os.Exit(1)
		}
		time.Sleep(time.Second * time.Duration(reconnectSec))
	}
	Handlers := handler.NewHandlers(redis)

	router := mux.NewRouter()
	Handlers.AddRoutes(router)
	log.Printf("Listening to 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
