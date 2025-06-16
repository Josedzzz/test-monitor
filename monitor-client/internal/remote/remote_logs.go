package remote

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type LogPayload struct {
	VMID    string `json:"vm_id"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

// sends a json log message to the central server via http post
func SendLog(serverURL, vmID, level, message string) error {
	payload := LogPayload{
		VMID:    vmID,
		Level:   level,
		Message: message,
		Time:    time.Now().Format(time.RFC3339),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = http.Post(serverURL+"/logs", "application/json", bytes.NewBuffer(data))
	return err
}
