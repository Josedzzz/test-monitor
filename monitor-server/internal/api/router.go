package api

import (
	"github.com/Josedzzz/monitor-server/internal/handlers"
	"github.com/gorilla/mux"
)

// setup the http routes for the monitoring server
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/logs", handlers.ReceiveLog).Methods("POST")
	r.HandleFunc("/machines", handlers.ListMachines).Methods("GET")
	r.HandleFunc("/logs/${vm_id}", handlers.GetLogsByVm).Methods("GET")
	return r
}
