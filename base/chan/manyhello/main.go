package main

import (
	"fmt"
)

func main() {
	done := make(chan int, 2)

	for i := 0; i < cap(done); i++ {
		go func(i int) {
			fmt.Println("hello world!", i)
			done <- 1
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		<-done
		fmt.Println("<-done")
	}
}
