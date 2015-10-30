package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	onceFunc := func() {
		fmt.Println("Only once")
	}

	once.Do(onceFunc)
	once.Do(onceFunc)
}
