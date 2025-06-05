package logs

import (
	"log"
	"os"
	"time"
)

var logger log.Logger

func InitLog() error {
	f, err := os.OpenFile("logs/monitor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	logger = *log.New(f, "", log.LstdFlags)
	return nil
}

func Info(msg string) {
	logger.Println(time.Now().Format(time.RFC3339) + " INFO: " + msg)
}
