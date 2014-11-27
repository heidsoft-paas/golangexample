package main

import (
        "fmt"
	"bytes"
	"bufio"
	"os"
        "github.com/fsouza/go-dockerclient"
)
//ref https://github.com/fsouza/go-dockerclient

func main() {
	createOpts := docker.CreateExecOptions{
		//Container:    "d0655935ca3ceb01af621e871e05c3d03bd533c2132562ac1622cd3a8ae2730d",
		Container:    "test1",
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
