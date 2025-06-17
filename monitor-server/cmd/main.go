package main

import (
	"log"
	"net/http"

	"github.com/Josedzzz/monitor-server/internal/api"
)

func main() {
	router := api.NewRouter()

	log.Println("Log server running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
