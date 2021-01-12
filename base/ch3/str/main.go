package main

import "fmt"

func main() {
	s := "hello, 世界"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println([]rune(s))
	fmt.Println(string(s[7:10]))
}
