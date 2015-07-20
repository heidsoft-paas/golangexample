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

	conn, _ := rados.NewConn()
        conn.ReadDefaultConfigFile()
        conn.Connect()
	defer conn.Shutdown()

        poolname := "pool1"
	imgName := "user1_image2"
	//snapName := "user1_image1_snap2"

        ioctx, err := conn.OpenIOContext(poolname)
        if err != nil {
		fmt.Printf("open pool failed:  %v", err)
		return
	}

	defer ioctx.Destroy()
	
	img := rbd.GetImage(ioctx, imgName)
	/*if err = img.Open(snapName); err != nil{
		fmt.Printf("open image %s with snap %s failed: %v", imgName, snapName, err)
	}*/
	if err = img.Open(); err != nil{
		fmt.Printf("open image %s failed: %v", imgName, err)
	}

	defer img.Close()
	
	if info, err := img.Stat(); err != nil{
		fmt.Printf("stat failed %v\n", err)
	}else{
		fmt.Printf("%#v\n", info)
	}

	parentPool := make([]byte, 128)
	parentName := make([]byte, 128)
	parentSnapname := make([]byte, 128)
	if err = img.GetParentInfo(parentPool, parentName, parentSnapname); err != nil {
		fmt.Printf("Get parent info failed: %v ", err)
	}else{
		fmt.Printf("pool:%s, pname:%s, psnapname:%s\n", string(parentPool[:]), string(parentName[:]), string(parentSnapname[:]))
	}
}
