package main

import "fmt"
import "strings"

func main(){
    i := 0xff
    fmt.Println(i)
    s := "1,5,3,4"
    set := strings.Split(s, ",")
    fmt.Printf("%s\n", set[len(set)-1])
}
