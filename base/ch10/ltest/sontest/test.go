package ltest

import (
	"flag"
	"fmt"
	"os"
)

var num = flag.Int("num", 1, "num....")

// PrintNumArgs 输出num参数
func PrintNumArgs() {
	flag.Parse()
	fmt.Print("ddd:", *num)
	fmt.Print(os.Args)
}
