package instances

import (
	"github.com/Oleg-Smal-git/diploma/services/physics/ecs"
)

// SystemRegistrar wraps all systems to be queried during preemptive memory allocation.
// Notice how the order in which systems are registered here is in fact the order in which
// they are going to be executed later on during the actual simulation.
var SystemRegistrar = []ecs.System{
	SystemMover{},
	SystemCollider{},
}

// Confirm that system structures satisfy System interface.
// This will throw a compile error otherwise.
var (
	_ ecs.System = (*SystemMover)(nil)
	_ ecs.System = (*SystemCollider)(nil)
)

// SystemMover implements the movement logic for Entity objects with ComponentPosition and ComponentVelocity.
type SystemMover struct{}

// Archetype returns a minimal required bitset for the system.
func (SystemMover) Archetype() ecs.ComponentID {
	return ComponentIDActive | ComponentIDPosition | ComponentIDVelocity
}

// Run performs one atomic step of the system logic.
func (SystemMover) Run(entity *ecs.Entity, entities *[]ecs.Entity) {
	// TODO: implement
	panic("")
}

// SystemCollider implements the collision logic for Entity objects with ComponentPosition and ComponentVelocity.
type SystemCollider struct{}

// Archetype returns a minimal required bitset for the system.
func (SystemCollider) Archetype() ecs.ComponentID {
	return ComponentIDActive | ComponentIDRigidBody | ComponentIDPosition | ComponentIDVelocity
}

// Run performs one atomic step of the system logic.
func (SystemCollider) Run(entity *ecs.Entity, entities *[]ecs.Entity) {
	// TODO: implement
	panic("")
}
