package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/dockerclient"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

func StopContainer(cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		err := dockerclient.StopContainer(cli, id)
		if err != nil {
			http.Error(w, "Error stopping container: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"message":      "Container stopped successfully",
			"container_id": id,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

