package dockerclient

import (
	"context"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// struct for the docker container info
type ContainerInfo struct {
	Id     string   `json:"id"`
	Names  []string `json:"names"`
	Image  string   `json:"image"`
	State  string   `json:"state"`
	Status string   `json:"status"`
}

// returns a list of all containers, running and stopped
func ListContainers() ([]ContainerInfo, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	var result []ContainerInfo
	for _, c := range containers {
		result = append(result, ContainerInfo{
			Id:     c.ID[:12],
			Names:  c.Names,
			Image:  c.Image,
			State:  c.State,
			Status: c.Status,
		})
	}
	return result, nil
}

// starts a container by the id using the given docker client
func StartContainer(cli *client.Client, containerID string) error {
	return cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
}

// stops a container by the id using the given docker client
func StopContainer(cli *client.Client, containerID string) error {
	return cli.ContainerStop(context.Background(), containerID, nil)
}

// inspects the info of a container by the id using the given docker client
func InspectContainer(cli *client.Client, containerID string) (types.ContainerJSON, error) {
	return cli.ContainerInspect(context.Background(), containerID)
}

// gets the container logs by the id and a tail
func GetContainerLogs(cli *client.Client, containerID string, tail string) (string, error) {
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Tail:       tail,
	}
	out, err := cli.ContainerLogs(context.Background(), containerID, options)
	if err != nil {
		return "", err
	}
	defer out.Close()
	var sb strings.Builder
	_, err = io.Copy(&sb, out)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
