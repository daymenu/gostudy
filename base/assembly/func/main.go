package main

import (
	"fmt"

	"github.com/daymenu/gostudy/base/assembly/func/pkg"
)

func main() {
	a, b := pkg.Swap(1, 2)
	fmt.Println(a, b)
}
