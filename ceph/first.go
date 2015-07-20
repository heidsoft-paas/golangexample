package main

import (
        "fmt"
        "github.com/noahdesu/go-ceph/rados"
        "github.com/noahdesu/go-ceph/rbd"
	"os/exec"
)

func GetUUID() string {
        out, _ := exec.Command("uuidgen").Output()
        return string(out[:36])
}

func main(){
	var ioctx *rados.IOContext

        createdList := []string{}
        
	conn, _ := rados.NewConn()
        conn.ReadDefaultConfigFile()
        conn.Connect()

	defer conn.Shutdown()

        poolname := "tmppool"
        err := conn.MakePool(poolname)
        if err != nil {
		fmt.Println("make pool failed:  %v", err)
		return
	}

	defer conn.DeletePool(poolname)

        ioctx, err = conn.OpenIOContext(poolname)
        if err != nil {
		fmt.Println("open pool failed:  %v", err)
		return
	}
	
	defer ioctx.Destroy()
        for i := 0; i < 10; i++ {
                name := GetUUID()
                _, err := rbd.Create(ioctx, name, 1<<22)
        	if err != nil {
			fmt.Println("make pool failed:  %v", err)
			return			
		}
                createdList = append(createdList, name)
        }

        //imageNames, err := rbd.GetImageNames(ioctx)

        for _, name := range createdList {
                img := rbd.GetImage(ioctx, name)
		fmt.Println("%v", img)
                img.Remove()
        }
}
