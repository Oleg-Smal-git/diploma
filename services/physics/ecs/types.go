package ecs

import (
	"github.com/Oleg-Smal-git/diploma/services/physics/runner"
)

type (
	// Chunk is a collection of Entity objects that have a certain set of Component flags.
	// Entities within same chunk are referenced to as archetype.
	Chunk struct {
		// Archetype is the bitset of ComponentID flags that Entity owns.
		Archetype ComponentID
		// Entities is a collection of Entity objects.
		Entities []Entity
		// Systems is a collection of System objects that are satisfied with parent Chunk's Archetype.
		Systems []System
	}

	// Entity is effectively a collection of Component objects.
	Entity struct {
		// Components is a map of Component objects that describe the state of an Entity.
		Components map[ComponentID]Component
	}

	// Component represents a flag that System objects use when iterating over Entity collections
	// in order to only alter / interact with entities that are meant to for respective systems,
	// and information relevant for those systems.
	Component interface {
		// ID returns the ComponentID of the parent Component.
		ID() ComponentID
		// New allocates all the required memory for the Component.
		New() Component
	}

	// ComponentID is an alias for uint8.
	ComponentID uint8

	// System contains the business logic of exactly ONE simulation rule.
	// The specific implementations of System interface also contain
	// buffers with pre-allocated memory for Component queries.
	System interface {
		// Archetype returns a minimal required bitset for the system.
		Archetype() ComponentID
		// Run performs one atomic step of the system logic.
		Run(*Entity, *[]Entity)
		// New allocates all the required memory for the System.
		New() System
		// Restore propagates simulation globals to child System objects.
		Restore(runner.Globals)
	}
)
