package main

import (
	"fmt"
	"sync"
)

type age int

func (n age) Compare(a interface{}) bool {
	a1, ok := a.(age)
	if !ok {
		return false
	}
	return n >= a1
}

var list *LinkedList = NewList()

func main() {
	fmt.Println("linked_list")

	var a age = 30
	var a1 age = 31

	fmt.Printf("列表是否为空：%t\n", list.Empty())
	list.Add(a)
	list.Print()
	fmt.Printf("列表长度：%d\n", list.Len())
	list.Add(a1)
	list.Print()
	fmt.Printf("列表长度：%d\n", list.Len())
	list.Delete(a)
	list.Print()
	fmt.Printf("列表长度：%d\n", list.Len())

	wait := sync.WaitGroup{}
	wait.Add(2)
	go func() {
		var a age = 55
		list.Add(a)
		wait.Done()
	}()

	go func() {
		var a age = 66
		list.Add(a)
		wait.Done()
	}()

	wait.Wait()
	list.Print()
}
