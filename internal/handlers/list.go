package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/dockerclient"
)

func GetContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := dockerclient.ListContainers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}
