package bank

import (
	"fmt"
	"testing"
	"time"
)

func TestTransfer(t *testing.T) {
	go func() {
		Deposit(200)
		fmt.Println(Balance())
	}()

	go Deposit(100)

	time.Sleep(1 * time.Second)
	fmt.Println(Balance())
	t.Fatal()
}

func TestMap(t *testing.T) {
	var x []int

	go func() { x = make([]int, 1000000) }()
	go func() { x = make([]int, 10) }()
	time.Sleep(1 * time.Second)
	x[999999] = 1
}
