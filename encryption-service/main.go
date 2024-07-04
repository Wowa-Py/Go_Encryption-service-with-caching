package main

import (
	"encryption-service/config"
	"encryption-service/handlers"
	"encryption-service/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)

	r.HandleFunc("/encrypt", handlers.EncryptHandler(cfg)).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
