package main

import (
	"fmt"
	"sort"
	"unsafe"
)

var a = []float64{4, 2, 5, 7, 2, 1, 88, 1.1}

func main() {
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	sort.Ints(b)
	fmt.Println(int(b[0]))
}
