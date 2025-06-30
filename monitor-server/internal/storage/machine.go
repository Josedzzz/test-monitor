package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Josedzzz/monitor-server/internal/models"
)

// called when a log is recived to update the vm info
func UpdateMachineInfo(vmId, ip, message string) {
	info := models.MachineInfo{
		VMID:     vmId,
		IP:       ip,
		LastSeen: time.Now(),
	}

	saveMachineInfoToDisk(info)
}

// returns the machines info
func GetAllMachines() []models.MachineInfo {
	dir := "machine_info"
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var result []models.MachineInfo
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err == nil {
				var m models.MachineInfo
				if json.Unmarshal(data, &m) == nil {
					result = append(result, m)
				}
			}
		}
	}
	return result
}

func saveMachineInfoToDisk(info models.MachineInfo) {
	dir := "machine_info"
	_ = os.MkdirAll(dir, os.ModePerm)

	path := filepath.Join(dir, fmt.Sprintf("%s.json", info.VMID))
	data, _ := json.MarshalIndent(info, "", "  ")
	_ = os.WriteFile(path, data, 0o644)
}
