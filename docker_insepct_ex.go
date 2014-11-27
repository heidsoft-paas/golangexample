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

	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	container, err := client.InspectContainer(name)
	if err != nil {
		fmt.Println("failed to inspect container %v", err)
		os.Exit(1)
	}
	state := container.State
	fmt.Println("ID: ", container.ID)
	fmt.Println("Pid: ", state.Pid)
}
