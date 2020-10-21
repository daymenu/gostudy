package ltest

import "testing"

func TestPrintNumArgs(t *testing.T) {
	PrintNumArgs()
	t.Error("错误了")
}
