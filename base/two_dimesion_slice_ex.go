package main

import "fmt"

func main(){
    var a [][]int
    b := [][]int{
        []int{0, 5, 12, 17},
        []int{6, 11, 18, 23},
    }
    
    a = b
    fmt.Println(a)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
