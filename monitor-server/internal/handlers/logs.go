package handlers

import (
	"encoding/json"
	"net"
	"net/http"
	"os"

	"github.com/Josedzzz/monitor-server/internal/models"
	"github.com/Josedzzz/monitor-server/internal/storage"
	"github.com/gorilla/mux"
)

// returns the log content for a specific vm
func GetLogsByVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmID := vars["vm_id"]
	logPath := "logs/" + vmID + ".log"

	data, err := os.ReadFile(logPath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Log file not found for VM: " + vmID,
			"error":   err.Error(),
		})
		return
	}

	response := map[string]string{
		"message": "Log file retrieved successfully",
		"log":     string(data),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handles incoming log data from a monitor-client instance
func ReceiveLog(w http.ResponseWriter, r *http.Request) {
	var payload models.LogPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ip := getIp(r.RemoteAddr)
	storage.UpdateMachineInfo(payload.VMID, ip, payload.Message)

	if err := storage.SaveLog(payload); err != nil {
		http.Error(w, "Failed to save log", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// extracts the ip address from request
func getIp(remoteAddr string) string {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return host
}
