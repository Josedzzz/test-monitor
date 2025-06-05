package monitor

import (
	"context"
	"time"

	"github.com/Josedzzz/test-monitor/internal/logs"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func StartMonitoring(cli *client.Client, interval time.Duration) {
	go func() {
		for {
			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
			if err != nil {
				logs.Info("Error listing containers: " + err.Error())
			} else {
				for _, c := range containers {
					state := c.State
					name := c.Names[0]
					logs.Info("Container: " + name + " is " + state)
				}
			}
			time.Sleep(interval)
		}
	}()
}
