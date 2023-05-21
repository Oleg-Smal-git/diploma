//go:build BUILD_ECS

package main

import (
	"github.com/Oleg-Smal-git/diploma/config"
	"github.com/Oleg-Smal-git/diploma/services/archivist"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
	"github.com/Oleg-Smal-git/diploma/services/physics/ecs"
	"github.com/Oleg-Smal-git/diploma/services/physics/ecs/instances/turbulence"
)

func initialize() (interfaces.Runner, interfaces.Archivist) {
	return ecs.NewRunner(
		turbulence.NewStater(), turbulence.ComponentRegistrar,
		turbulence.ArchetypesRegistrar, turbulence.SystemRegistrar,
		config.StateCapacity,
	), archivist.NewArchivist(config.MarshalFunctor, config.UnmarshalFunctor)
}
