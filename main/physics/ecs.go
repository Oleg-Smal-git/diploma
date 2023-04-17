//go:build BUILD_ECS

package main

import (
	"github.com/Oleg-Smal-git/diploma/services/archivist"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
	"github.com/Oleg-Smal-git/diploma/services/physics/ecs"
	"github.com/Oleg-Smal-git/diploma/services/physics/ecs/registrar"
)

func initialize() (interfaces.Runner, interfaces.Archivist) {
	return ecs.NewRunner(registrar.ComponentRegistrar, registrar.ArchetypesRegistrar, registrar.SystemRegistrar), archivist.NewArchivist()
}
