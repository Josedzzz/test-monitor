package handlers

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/Josedzzz/monitor-server/internal/models"
	"github.com/Josedzzz/monitor-server/internal/storage"
)

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
