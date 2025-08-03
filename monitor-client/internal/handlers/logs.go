package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/dockerclient"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

// handles the request to get the container logs
func GetContainerLogsHandler(cli *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		containerID := vars["id"]
		tail := r.URL.Query().Get("tail")
		if tail == "" {
			tail = "100"
		}
		logs, err := dockerclient.GetContainerLogs(cli, containerID, tail)
		if err != nil {
			http.Error(w, "Failed to get logs: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"logs": logs})
	}
}
