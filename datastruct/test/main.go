package main

import "fmt"

var aa = cc
var cc = 1

func main() {
	var str string = "hello world"
	fmt.Printf("%s %[1]T", str)
}
