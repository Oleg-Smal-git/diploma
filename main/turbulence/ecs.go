//go:build BUILD_ECS

package main

import (
	"runtime/debug"

	"github.com/Oleg-Smal-git/diploma/config"
	"github.com/Oleg-Smal-git/diploma/services/archivist"
	"github.com/Oleg-Smal-git/diploma/services/ecs"
	instances "github.com/Oleg-Smal-git/diploma/services/instances/ecs"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

func initialize() (interfaces.Runner, interfaces.Archivist) {
	debug.SetGCPercent(0) // Disable automatic garbage collection.
	return ecs.NewRunner(
		instances.NewStater(), instances.ComponentRegistrar,
		instances.ArchetypesRegistrar, instances.SystemRegistrar,
		config.StateCapacity,
	), archivist.NewArchivist(config.MarshalFunctor, config.UnmarshalFunctor)
}
