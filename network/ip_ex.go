package main

import "net"
import "fmt"

func main(){
    var(
        ip net.IP
        network *net.IPNet
	)
    mask := net.CIDRMask(16, 32)
    fmt.Println(mask.String())
    
    ip, network, _ = net.ParseCIDR("192.168.1.10/24");
    fmt.Println(ip.String())
    fmt.Println(network.IP.String())
    fmt.Println(network.Mask.String())
}
