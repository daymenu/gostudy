package main

import "github.com/google/wire"

// Foo foo
type Foo int

// Bar bar
type Bar int

// ProvideFoo provide foo
func ProvideFoo() Foo {
	return 1
}

// ProvideBar provide bar
func ProvideBar() Bar {
	return 2
}

// FooBar foobar
type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

// Set set
var Set = wire.NewSet(
	ProvideBar,
	ProvideFoo,
	// wire.Struct(new(FooBar), "MyFoo", "MyBar"),
	wire.Struct(new(FooBar), "*"), // *代表全部字段
)
