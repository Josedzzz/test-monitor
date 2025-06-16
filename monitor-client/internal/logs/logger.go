package logs

import (
	"log"
	"os"
	"time"
)

var logger log.Logger

// creates or opens the log file and sets ups the logger
func InitLog() error {
	f, err := os.OpenFile("logs/monitor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	logger = *log.New(f, "", log.LstdFlags)
	return nil
}

// writes a log entry with the time and a msg
func Info(msg string) {
	logger.Println(time.Now().Format(time.RFC3339) + " INFO: " + msg)
}
