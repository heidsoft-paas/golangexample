package main

import (
        "fmt"
        "github.com/fsouza/go-dockerclient"
)
//ref https://github.com/fsouza/go-dockerclient

func main() {
        endpoint := "unix:///var/run/docker.sock"
        client, _ := docker.NewClient(endpoint)
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})
	for _, container := range containers {
		fmt.Println("ID: ", container.ID)
		fmt.Println("Img: ", container.Image)
		fmt.Println("Command: ", container.Command)
		fmt.Println("Created: ", container.Created)
		fmt.Println("Status: ", container.Status)
		fmt.Println("Port: ", container.Ports)
		fmt.Println("SizeRw: ", container.SizeRw)
		fmt.Println("SizeRootFs: ", container.SizeRootFs)
		fmt.Println("Names: ", container.Names)
	}

}
