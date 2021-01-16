package main

import (
	"fmt"
)

func main() {
	var a [3]int = [3]int{1, 2, 3}
	b := &a
	fmt.Printf("%T %T\n", a, b)
	//[3]int *[3]int

	fmt.Println(a[0], a[1], a[2]) //打印数组a的元素
	// 1 2 3

	fmt.Println(b[0], b[1], b[2]) // 通过数组指针访问数组元素的方式和数组类似
	// 1 2 3

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

}
