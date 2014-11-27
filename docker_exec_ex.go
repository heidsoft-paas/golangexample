package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

//ref https://github.com/fsouza/go-dockerclient

func main() {
	var container string
	flag.StringVar(&container, "c", "test1", "container name")

	flag.Parse()

	createOpts := docker.CreateExecOptions{
		Container:    container,
		Cmd:          []string{"ifconfig"},
		AttachStdin:  false,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
	}
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	execObj, err := client.CreateExec(createOpts)
	if err != nil {
		fmt.Println("failed to run in container - Exec setup failed - %v", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	wrBuf := bufio.NewWriter(&buf)
	startOpts := docker.StartExecOptions{
		Detach:       false,
		Tty:          false,
		OutputStream: wrBuf,
		ErrorStream:  wrBuf,
		RawTerminal:  false,
	}
	errChan := make(chan error, 1)
	go func() {
		errChan <- client.StartExec(execObj.ID, startOpts)
	}()
	err = <-errChan
	if err != nil {
		fmt.Println("failed to run in container - Exec start failed - %v", err)
		os.Exit(1)
	}
	wrBuf.Flush()
	//data := buf.Bytes()
	fmt.Println("result: ", buf.String())
}
