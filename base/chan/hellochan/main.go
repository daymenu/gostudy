package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		fmt.Println("你好，世界！")
		ch <- 1
	}()

	<-ch
}
