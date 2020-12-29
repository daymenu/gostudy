//+build wireinject

package main

import "github.com/google/wire"

func injectFooBar() FooBar {
	wire.Build(Set)
	return FooBar{}
}
