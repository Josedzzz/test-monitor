package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Josedzzz/test-monitor/internal/api"
	"github.com/Josedzzz/test-monitor/internal/logs"
	"github.com/Josedzzz/test-monitor/internal/monitor"
	"github.com/docker/docker/client"
)

func NewDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func main() {
	cli, err := NewDockerClient()
	if err != nil {
		log.Fatalf("Error while creating the docker client: %v", err)
	}
	r := api.NewRouter(cli)
	_ = logs.InitLog()
	monitor.StartMonitoring(cli, 30*time.Minute)
	fmt.Println("Server is up!!!")
	log.Fatal(http.ListenAndServe("0.0.0.0:81", r))
}
