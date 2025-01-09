// go:build wireinject
//go:build wireinject
// +build wireinject

// The build tag makes sure this file is not in the final build

package main

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/fsufitch/halebopp-hotspot/cmd"
	"github.com/fsufitch/halebopp-hotspot/impl/dummy"
	"github.com/google/wire"
)

func main() {
	cmd.Entrypoint(initializeDummyHaleBopp)
}

func initializeDummyHaleBopp() (*halebopp.HaleBopp, func(), error) {
	panic(wire.Build(
		halebopp.ProvideHaleBopp,
		dummy.ProvideDummyImplementations,
	))
}
