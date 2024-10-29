package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type MyContainer struct {
	ID     string
	Image  string
	Status string
}

func listContainers() ([]MyContainer, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}

	filterArgs := filters.NewArgs()
	filterArgs.Add("status", "running")

	options := types.ContainerListOptions{
		All:     false,
		Filters: filterArgs,
	}

	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	var myContainers []MyContainer
	for _, container := range containers {
		myContainers = append(myContainers, MyContainer{
			ID:     container.ID,
			Image:  container.Image,
			Status: container.Status,
		})
	}

	return myContainers, nil
}

func main() {
	containers, err := listContainers()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, c := range containers {
		fmt.Printf("Container ID: %s, Image: %s, Status: %s\n", c.ID, c.Image, c.Status)
	}
}
