//+build wireinject

package main

import (
	"os"

	"github.com/google/wire"
)

func injectFile() (*os.File, func(), error) {
	panic(wire.Build(provideFile))
}
