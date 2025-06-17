package models

// represents a log message sent by a VM
type LogPayload struct {
	VMID    string `json:"vm_id"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Time    string `json:"time"`
}
