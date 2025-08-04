package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/dockerclient"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

// handles get request to return the stats of a container
func GetContainerStatsHandler(cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		stats, err := dockerclient.GetContainerStats(cli, id)
		if err != nil {
			http.Error(w, "Error retrieving stats: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stats)
	}
}
