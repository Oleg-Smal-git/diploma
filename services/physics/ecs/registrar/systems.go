package registrar

import (
	"math"

	"github.com/Oleg-Smal-git/diploma/services/physics/ecs"
	"github.com/Oleg-Smal-git/diploma/services/physics/runner"
)

// SystemRegistrar wraps all systems to be queried during preemptive memory allocation.
// Notice how the order in which systems are registered here is in fact the order in which
// they are going to be executed later on during the actual simulation.
var SystemRegistrar = []ecs.System{
	&SystemMover{},
	&SystemCollider{},
	&SystemBoundary{},
}

// Confirm that system structures satisfy System interface.
// This will throw a compile error otherwise.
var (
	_ ecs.System = (*SystemMover)(nil)
	_ ecs.System = (*SystemCollider)(nil)
	_ ecs.System = (*SystemBoundary)(nil)
)

// SystemMover implements the movement logic.
type SystemMover struct {
	globals runner.Globals

	active   *ComponentActive
	position *ComponentPosition
	velocity *ComponentVelocity
}

// Archetype returns a minimal required bitset for the system.
func (SystemMover) Archetype() ecs.ComponentID {
	return ComponentIDActive | ComponentIDPosition | ComponentIDVelocity
}

// Run performs one atomic step of the system logic.
func (s *SystemMover) Run(entity *ecs.Entity, entities *[]ecs.Entity) {
	// Assert component types and cache them in buffer.
	if s.active = entity.Components[ComponentIDActive].(*ComponentActive); !s.active.Active {
		return
	}
	s.position = entity.Components[ComponentIDPosition].(*ComponentPosition)
	s.velocity = entity.Components[ComponentIDVelocity].(*ComponentVelocity)

	// Business logic.
	s.position.X += s.velocity.X * s.globals.FrameDuration
	s.position.Y += s.velocity.Y * s.globals.FrameDuration
}

// New allocates all the required memory for the System.
func (SystemMover) New() ecs.System {
	// It's okay to keep pointers as nil since those will just be redirected
	// to already existing components during Run calls.
	return &SystemMover{}
}

// Restore propagates simulation globals to child System objects.
func (s *SystemMover) Restore(globals runner.Globals) {
	s.globals = globals
}

// SystemCollider implements the collision logic.
type SystemCollider struct {
	globals runner.Globals

	active    *ComponentActive
	rigidBody *ComponentRigidBody
	position  *ComponentPosition
	velocity  *ComponentVelocity
}

// Archetype returns a minimal required bitset for the system.
func (SystemCollider) Archetype() ecs.ComponentID {
	return ComponentIDActive | ComponentIDRigidBody | ComponentIDPosition | ComponentIDVelocity
}

// Run performs one atomic step of the system logic.
func (s *SystemCollider) Run(entity *ecs.Entity, entities *[]ecs.Entity) {
	// Assert component types and cache them in buffer.
	if s.active = entity.Components[ComponentIDActive].(*ComponentActive); !s.active.Active {
		return
	}
	s.rigidBody = entity.Components[ComponentIDRigidBody].(*ComponentRigidBody)
	s.position = entity.Components[ComponentIDPosition].(*ComponentPosition)
	s.velocity = entity.Components[ComponentIDVelocity].(*ComponentVelocity)

	// Business logic.

}

// New allocates all the required memory for the System.
func (SystemCollider) New() ecs.System {
	// It's okay to keep pointers as nil since those will just be redirected
	// to already existing components during Run calls.
	return &SystemCollider{}
}

// Restore propagates simulation globals to child System objects.
func (s *SystemCollider) Restore(globals runner.Globals) {
	s.globals = globals
}

// SystemBoundary implements the boundary logic.
type SystemBoundary struct {
	globals runner.Globals

	active    *ComponentActive
	rigidBody *ComponentRigidBody
	boundary  *ComponentBoundary
	position  *ComponentPosition
	velocity  *ComponentVelocity

	delta float64
}

// Archetype returns a minimal required bitset for the system.
func (SystemBoundary) Archetype() ecs.ComponentID {
	return ComponentIDActive | ComponentIDRigidBody | ComponentIDBoundary | ComponentIDPosition | ComponentIDVelocity
}

// Run performs one atomic step of the system logic.
func (s *SystemBoundary) Run(entity *ecs.Entity, entities *[]ecs.Entity) {
	// Assert component types and cache them in buffer.
	if s.active = entity.Components[ComponentIDActive].(*ComponentActive); !s.active.Active {
		return
	}
	s.rigidBody = entity.Components[ComponentIDRigidBody].(*ComponentRigidBody)
	s.boundary = entity.Components[ComponentIDBoundary].(*ComponentBoundary)
	s.position = entity.Components[ComponentIDPosition].(*ComponentPosition)
	s.velocity = entity.Components[ComponentIDVelocity].(*ComponentVelocity)

	// Business Logic.
	if s.delta = s.position.X + s.rigidBody.Size - s.boundary.MaxX; s.delta > 0 {
		s.position.X -= s.delta // Right edge.
		s.velocity.X = -math.Abs(s.velocity.X)
	}
	if s.delta = s.position.X - s.rigidBody.Size - s.boundary.MinX; s.delta < 0 {
		s.position.X -= s.delta // Left edge.
		s.velocity.X = math.Abs(s.velocity.X)
	}
	if s.delta = s.position.Y + s.rigidBody.Size - s.boundary.MaxY; s.delta > 0 {
		s.position.Y -= s.delta // Bottom edge.
		s.velocity.Y *= -math.Abs(s.velocity.Y)
	}
	if s.delta = s.position.Y - s.rigidBody.Size - s.boundary.MinY; s.delta < 0 {
		s.position.Y -= s.delta // Top edge.
		s.velocity.Y *= math.Abs(s.velocity.Y)
	}
}

// New allocates all the required memory for the System.
func (SystemBoundary) New() ecs.System {
	// It's okay to keep pointers as nil since those will just be redirected
	// to already existing components during Run calls.
	return &SystemBoundary{}
}

// Restore propagates simulation globals to child System objects.
func (s *SystemBoundary) Restore(globals runner.Globals) {
	s.globals = globals
}
