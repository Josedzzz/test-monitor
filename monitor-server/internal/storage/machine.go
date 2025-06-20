package storage

import (
	"strings"
	"sync"
	"time"

	"github.com/Josedzzz/monitor-server/internal/models"
)

var (
	machineMap = make(map[string]models.MachineInfo)
	mu         sync.Mutex
)

// called when a log is recived to update the vm info
func UpdateMachineInfo(vmId, ip, message string) {
	mu.Lock()
	defer mu.Unlock()

	containers := extractContainerFromMessage(message)

	machineMap[vmId] = models.MachineInfo{
		VMID:       vmId,
		IP:         ip,
		LastSeen:   time.Now(),
		Containers: containers,
	}
}

// returns the machines info
func GetAllMachines() []models.MachineInfo {
	mu.Lock()
	defer mu.Unlock()

	var list []models.MachineInfo
	for _, m := range machineMap {
		list = append(list, m)
	}
	return list
}

// container name extraction (example: "Container: /nginx is running")
func extractContainerFromMessage(msg string) []string {
	if strings.Contains(msg, "Container: ") {
		parts := strings.Split(msg, " ")
		if len(parts) >= 2 {
			return []string{parts[1]}
		}
	}
	return nil
}
