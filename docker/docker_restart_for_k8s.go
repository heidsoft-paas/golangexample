package main

import (
	"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

//ref https://github.com/fsouza/go-dockerclient

func main() {
	var(
        ID string
        NetContainer string
    )
	flag.StringVar(&ID, "id", "", "container id")
    flag.StringVar(&NetContainer, "net", "", "network container id")

	flag.Parse()

    if ID == "" || NetContainer == ""{
        fmt.Println("container id or network container id is null")
        os.Exit(1)
    }
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)

	err := client.StartContainer(ID, &docker.HostConfig{
		NetworkMode:  "container:"+NetContainer,
	})
	if err != nil {
		fmt.Println("failed to start container, %v", err)
		os.Exit(1)
	}
}
