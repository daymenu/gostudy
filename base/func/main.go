package main

import "fmt"

func main() {
	var a = 3
	a1 := func() int {
		return a
	}()
	a++
	fmt.Println(a, a1)
	fmt.Println(&a, &a1)
}
