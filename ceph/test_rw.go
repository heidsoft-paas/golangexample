package main

import (
        "fmt"
        "github.com/noahdesu/go-ceph/rados"
	"os/exec"
)

func GetUUID() string {
        out, _ := exec.Command("uuidgen").Output()
        return string(out[:36])
}

func main(){
	var ioctx *rados.IOContext

	conn, _ := rados.NewConn()
        conn.ReadDefaultConfigFile()
        conn.Connect()
	defer conn.Shutdown()

        poolname := "tmppool"

        ioctx, err := conn.OpenIOContext(poolname)
        if err != nil {
		fmt.Printf("open pool failed:  %v", err)
		return
	}
	defer ioctx.Destroy()

	oid := "object1"
	data := make([]byte, 16)
	if dataLen, err := ioctx.Read(oid, data, 0); err != nil{
		if err == rados.RadosErrorNotFound {
			fmt.Printf("%s not found\n", oid)
		} else {
			fmt.Printf("read %s failed: %v", oid, err)
		}
	} else {
		fmt.Printf("read data len %d\n", dataLen)
	}
}
