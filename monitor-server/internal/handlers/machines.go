package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/monitor-server/internal/storage"
)

// returns information about all registered monitor-client instances
func ListMachines(w http.ResponseWriter, r *http.Request) {
	machines := storage.GetAllMachines()

	if machines == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Failed to fetch machine data",
		})
		return
	}

	response := map[string]any{
		"message":  "Machines fetched successfully",
		"count":    len(machines),
		"machines": machines,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
