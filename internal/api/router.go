package api

import (
	"github.com/Josedzzz/test-monitor/internal/handlers"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

func NewRouter(cli *client.Client) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/containers", handlers.GetContainers).Methods("GET")
	r.HandleFunc("/containers/{id}/start", handlers.StartContainer(cli)).Methods("POST")
	r.HandleFunc("/containers/{id}/stop", handlers.StopContainer(cli)).Methods("POST")
	return r
}
