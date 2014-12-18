package main
import "time"
import "fmt"
func main(){
    fmt.Println("start")
    go func(){
        for {
            time.Sleep(time.Second * 1)
            fmt.Println("do some work")
        }
    }()

    select{
    }
}
