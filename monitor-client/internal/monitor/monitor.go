package monitor

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Josedzzz/test-monitor/internal/logs"
	"github.com/Josedzzz/test-monitor/internal/remote"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

// launches a background goroutine that logs the state of all containers periodically
func StartMonitoring(cli *client.Client, interval time.Duration) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: could not load .env file, using system environment variables")
	}

	serverURL := os.Getenv("SERVER_URL")
	vmID := os.Getenv("VM_ID")

	if serverURL == "" || vmID == "" {
		log.Fatal("Missing SERVER_URL or VM_ID in environment")
	}

	go func() {
		for {
			// list all containers
			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
			if err != nil {
				msg := "Error listing containers: " + err.Error()
				log.Println(msg)
				_ = remote.SendLog(serverURL, vmID, "ERROR", msg)
			} else {
				// log the state of each container
				for _, c := range containers {
					state := c.State
					name := c.Names[0]
					msg := "Container: " + name + " is " + state
					logs.Info(msg)
					_ = remote.SendLog(serverURL, vmID, "INFO", msg)
				}
			}
			// wait for the next interval
			time.Sleep(interval)
		}
	}()
}
