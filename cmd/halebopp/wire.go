//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/fsufitch/halebopp-hotspot/impl/x278"
	"github.com/google/wire"
)

func initializeDefaultHaleBopp() (*halebopp.HaleBopp, func(), error) {
	panic(wire.Build(
		halebopp.ProvideHaleBopp,
		x278.ProvideX278,
	))
}
