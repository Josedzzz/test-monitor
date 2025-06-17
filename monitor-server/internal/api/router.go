package api

import (
	"github.com/Josedzzz/monitor-server/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/logs", handlers.ReceiveLog).Methods("POST")
	return r
}
