package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/daymenu/gostudy/base/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 36.7, "体温")

func main() {

	var w io.Writer
	w = os.Stdout
	if w == nil {
		fmt.Println("w is nil")
	}
	fmt.Printf("%T %[1]q\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T %[1]q\n", w)
	// flag.Parse()
	// fmt.Println(*temp)

	var buf *bytes.Buffer
	var w1 io.Writer = buf
	fmt.Printf("%T %[1]q\n", w1)
	if w1 == nil {
		fmt.Println("w is nil")
	}
	os.Stdout.WriteString()
	var rw io.ReadWriter
	rw = os.Stdout

	w2 := rw.(io.Writer)

	fmt.Printf("%T %[1]q\n", rw)
	fmt.Printf("%T %[1]q\n", w2)

}
