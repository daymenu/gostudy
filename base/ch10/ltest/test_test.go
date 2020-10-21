package ltest

import "testing"

func TestPrintNumArgs(t *testing.T) {
	PrintNumArgs()
	t.Fail()
}

func BenchmarkPrintNumArgs(b *testing.B) {
	b.Fail()
}
