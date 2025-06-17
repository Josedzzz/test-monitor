package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Josedzzz/monitor-server/internal/models"
	"github.com/Josedzzz/monitor-server/internal/storage"
)

func ReceiveLog(w http.ResponseWriter, r *http.Request) {
	var payload models.LogPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := storage.SaveLog(payload); err != nil {
		http.Error(w, "Failed to save log", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
