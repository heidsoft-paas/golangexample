package main

import (
	"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

// Restart container for kubernetes
// When RestartPolicy is "never", network container will be created newly.
// Add Pod's container to network container's network namespace.
// Ref https://github.com/fsouza/go-dockerclient

func main() {
	var(
        id string
        netContainer string
    )
	flag.StringVar(&id, "id", "", "container id")
    flag.StringVar(&netContainer, "net", "", "network container id")

	flag.Parse()

    if id == "" || netContainer == ""{
        fmt.Println("container id or network container id is null")
        os.Exit(1)
    }
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)

	err := client.StartContainer(id, &docker.HostConfig{
		NetworkMode:  "container:"+netContainer,
	})
	if err != nil {
		fmt.Println("failed to start container, %v", err)
		os.Exit(1)
	}
	fmt.Printf("start container %s successful", id)
}
