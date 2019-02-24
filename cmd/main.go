package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"redis-responder/handler"
	"redis-responder/storage"
)

func main() {
	host := os.Getenv("REDIS_HOST")
	redis := storage.NewRedis()
	if err := redis.Connect(host); err != nil {
		log.Printf("Cannot connect redis cluster (%s): %s\n", host, err)
		os.Exit(2)
	}
	Handlers := handler.NewHandlers(redis)

	router := mux.NewRouter()
	Handlers.AddRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
