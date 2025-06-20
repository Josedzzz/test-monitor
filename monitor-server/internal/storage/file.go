package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Josedzzz/monitor-server/internal/models"
)

// writes a log entry to a file named after the vmId and appends the log entry
func SaveLog(p models.LogPayload) error {
	logDir := "logs"

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return err
	}

	logFile := filepath.Join(logDir, fmt.Sprintf("%s.log", p.VMID))

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	entry := fmt.Sprintf("[%s] [%s]: %s\n", p.Time, p.Level, p.Message)
	_, err = f.WriteString(entry)
	return err
}
