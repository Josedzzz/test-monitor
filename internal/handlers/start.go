package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/dockerclient"
	"github.com/Josedzzz/test-monitor/internal/logs"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

// returns an http handler that strats a docker container by id
func StartContainer(cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		err := dockerclient.StartContainer(cli, id)
		if err != nil {
			http.Error(w, "Error starting container: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"message":      "Container started successfully",
			"container_id": id,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		logs.Info("Container " + id + " started")
	}
}
