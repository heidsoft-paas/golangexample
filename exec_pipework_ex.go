package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("pipework", "br1", "test", "172.16.213.190/16@172.16.213.2")
	cmd.Dir = "/usr/local/bin"
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("stderr: %q\n", out.String())
}

