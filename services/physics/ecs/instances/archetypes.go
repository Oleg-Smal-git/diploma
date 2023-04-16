package instances

import (
	"github.com/Oleg-Smal-git/diploma/services/physics/ecs"
)

// ArchetypesRegistrar contains all known archetypes of simulation objects.
var ArchetypesRegistrar = []ecs.ComponentID{
	ArchetypeBall,
}

// Archetypes are composites of ComponentID that exhaustively describe an Entity class.
const (
	// ArchetypeBall goes bounce :)
	ArchetypeBall = ComponentIDActive | ComponentIDRigidBody | ComponentIDPosition | ComponentIDVelocity
)
