package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/api"
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

	fmt.Println("Server is up!!!")
	log.Fatal(http.ListenAndServe("0.0.0.0:81", r))
}
