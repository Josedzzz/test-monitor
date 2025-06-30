package models

import "time"

type MachineInfo struct {
	VMID     string    `json:"vm_id"`
	IP       string    `json:"ip"`
	LastSeen time.Time `json:"last_seen"`
}
