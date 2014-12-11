package main

import (
	"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

//ref https://github.com/fsouza/go-dockerclient

func main() {
	var name string
	flag.StringVar(&name, "c", "test1", "container name")

	flag.Parse()

	opts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Cmd:          []string{"/bin/bash"},
			Hostname:     "host1",
			Image:        "ubuntu:14.04",
			Memory:       128000000,
			CPUShares:    2048,
			CPUSet:       "0,",
                	AttachStdin:  true,
                	AttachStdout: true,
                	AttachStderr: true,
			Tty: true,
		},
	}
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	dockerContainer, err := client.CreateContainer(opts)
	if err != nil {
		fmt.Println("failed to create container - Exec setup failed - %v", err)
		os.Exit(1)
	}

	err = client.StartContainer(dockerContainer.ID, &docker.HostConfig{
		NetworkMode:  "bridge",
	})
	if err != nil {
		fmt.Println("failed to start container - Exec setup failed - %v", err)
		os.Exit(1)
	}
}
