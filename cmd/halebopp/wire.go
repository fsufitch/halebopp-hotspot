//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/fsufitch/halebopp-hotspot/modules"
	"github.com/google/wire"
)

func initializeDefaultHaleBopp() (*halebopp.HaleBopp, func(), error) {
	panic(wire.Build(
		halebopp.ProvideHaleBopp,
		modules.ProvideX278Battery,
		modules.ProvideX278Charging,
	))
}
