package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Josedzzz/test-monitor/internal/api"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error while creating the docker client: %v", err)
	}

	r := api.NewRouter(cli)

	fmt.Println("Server in: http://localhost:8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:81", r))
}
