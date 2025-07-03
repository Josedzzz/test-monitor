package main

import (
	"log"
	"net/http"

	"github.com/Josedzzz/monitor-server/internal/api"
	"github.com/gorilla/handlers"
)

func main() {
	router := api.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)

	log.Println("Log server running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", corsHandler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
