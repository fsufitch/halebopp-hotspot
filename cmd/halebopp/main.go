// go:build wireinject
//go:build wireinject
// +build wireinject

// The build tag makes sure this file is not in the final build

package main

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/fsufitch/halebopp-hotspot/cmd"
	"github.com/fsufitch/halebopp-hotspot/impl/sim7600x"
	"github.com/fsufitch/halebopp-hotspot/impl/x278"
	"github.com/google/wire"
)

func main() {
	cmd.Entrypoint(initializeDefaultHaleBopp)
}

func initializeDefaultHaleBopp() (*halebopp.HaleBopp, func(), error) {
	panic(wire.Build(
		halebopp.ProvideHaleBopp,
		x278.ProvideX278,
		sim7600x.ProvideModem,
	))
}
