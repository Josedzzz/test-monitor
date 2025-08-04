package api

import (
	"github.com/Josedzzz/test-monitor/internal/handlers"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

// configures the http routes to their respective handlers
func NewRouter(cli *client.Client) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/containers", handlers.GetContainers).Methods("GET")
	r.HandleFunc("/containers/{id}/start", handlers.StartContainer(cli)).Methods("POST")
	r.HandleFunc("/containers/{id}/stop", handlers.StopContainer(cli)).Methods("POST")
	r.HandleFunc("/containers/{id}/inspect", handlers.InspectContainer(cli)).Methods("GET")
	r.HandleFunc("/containers/{id}/logs", handlers.GetContainerLogsHandler(cli)).Methods("GET")
	r.HandleFunc("/containers/{id}/stats", handlers.GetContainerStatsHandler(cli)).Methods("GET")
	return r
}
