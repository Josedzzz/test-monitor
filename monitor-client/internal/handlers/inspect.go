package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/dockerclient"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

// returns an http handler that inspects the container info
func InspectContainer(cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		info, err := dockerclient.InspectContainer(cli, id)
		if err != nil {
			http.Error(w, "Error inspecting the container: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	}
}
