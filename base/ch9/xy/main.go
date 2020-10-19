package main

import (
	"fmt"
	"time"
)

var x, y int

func main() {
	go func() {
		x = 1
		fmt.Println("y", y)
	}()
	go func() {
		y = 1
		fmt.Println("x", x)
	}()
	time.Sleep(1 * time.Second)

}
