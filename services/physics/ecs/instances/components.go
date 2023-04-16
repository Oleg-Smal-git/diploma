package instances

import (
	"github.com/Oleg-Smal-git/diploma/services/physics/ecs"
)

// ComponentRegistrar wraps all components to be queried during preemptive memory allocation.
var ComponentRegistrar = []ecs.Component{
	ComponentActive{},
	ComponentRigidBody{},
	ComponentPosition{},
	ComponentVelocity{},
}

// Components are stored in a bitset, which means that
// each one should be of kind 2^n where n ∈ N_0+.
const (
	// ComponentIDActive is used to simulate nullability. This allows the garbage collector to
	// allocate memory before the simulation starts to run, thus
	// significantly reducing the time that would have been spent on that otherwise.
	ComponentIDActive ecs.ComponentID = 1 << iota
	// ComponentIDRigidBody describes static object properties.
	ComponentIDRigidBody
	// ComponentIDPosition describes object position.
	ComponentIDPosition
	// ComponentIDVelocity describes object velocity.
	ComponentIDVelocity
)

// Confirm that all component structures satisfy Component interface.
// This will throw a compile error otherwise.
var (
	_ ecs.Component = (*ComponentActive)(nil)
	_ ecs.Component = (*ComponentRigidBody)(nil)
	_ ecs.Component = (*ComponentPosition)(nil)
	_ ecs.Component = (*ComponentVelocity)(nil)
)

// ComponentActive fakes nullability.
type ComponentActive struct {
	Active bool
}

// ID identifies the component.
func (c ComponentActive) ID() ecs.ComponentID {
	return ComponentIDActive
}

// ComponentRigidBody describes static object properties.
type ComponentRigidBody struct{}

// ID identifies the component.
func (c ComponentRigidBody) ID() ecs.ComponentID {
	return ComponentIDRigidBody
}

// ComponentPosition describes object position.
type ComponentPosition struct {
	X, Y float64
}

// ID identifies the component.
func (c ComponentPosition) ID() ecs.ComponentID {
	return ComponentIDPosition
}

// ComponentVelocity describes object velocity.
type ComponentVelocity struct {
	X, Y float64
}

// ID identifies the component.
func (c ComponentVelocity) ID() ecs.ComponentID {
	return ComponentIDVelocity
}
