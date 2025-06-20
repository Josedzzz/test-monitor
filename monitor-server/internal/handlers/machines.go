package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/monitor-server/internal/storage"
)

// returns information about all registered monitor-client instances
func ListMachines(w http.ResponseWriter, r *http.Request) {
	machines := storage.GetAllMachines()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(machines)
}
