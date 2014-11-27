package main

import (
	"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os/exec"
	"os"
	"bytes"
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
	
	//pid := strconv.Itoa(state.Pid)
	target := fmt.Sprintf("--target=%d", state.Pid)
	args := append([]string{}, target, "--net", "ifconfig")
	cmd := exec.Command("/usr/local/bin/nsenter", args...)
	cmd.Dir = "/usr/local/bin"
	fmt.Println("args:", cmd.Args)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		fmt.Println("run error: ", err)
		os.Exit(1)
	}
	fmt.Printf("run result: %q\n", out.String())
}
