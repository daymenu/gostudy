package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 1)

	ch <- 4
	go func() { fmt.Println(<-ch) }()
	time.Sleep(1 * time.Second)
	fmt.Println("hello")
}
