package ecs

import (
	"github.com/Oleg-Smal-git/diploma/services/physics/runner"
)

// Confirm that ECS satisfies runner.Runner interface.
// This will throw a compile error otherwise.
var _ runner.Runner = (*ECS)(nil)

// ECS stands for Entity Component System and is an architectural
// pattern we will be using for this implementation of runner.Runner.
type ECS struct {
	chunks []Chunk
}

// NewECS constructs an ECS object.
func NewECS(componentRegistrar []Component, archetypesRegistrar []ComponentID, systemRegistrar []System) ECS {
	ecs := ECS{
		chunks: make([]Chunk, 0, len(archetypesRegistrar)),
	}
	for _, a := range archetypesRegistrar {
		chunk := Chunk{
			Archetype: a,
			Entities:  make([]Entity, chunkCapacity),
			Systems:   make([]System, 0, len(systemRegistrar)),
		}
		for i := range chunk.Entities {
			chunk.Entities[i].Components = make(map[ComponentID]Component, len(componentRegistrar))
			for _, c := range componentRegistrar {
				if a&c.ID() == c.ID() {
					// This is done in order to deep copy the interface value.
					chunk.Entities[i].Components[c.ID()] = c.New()
				}
			}
		}
		for _, s := range systemRegistrar {
			if a&s.Archetype() == s.Archetype() {
				// This is done in order to deep copy the interface value.
				chunk.Systems = append(chunk.Systems, s.New())
			}
		}
		ecs.chunks[a] = chunk
	}
	return ecs
}

// Next performs one atomic step of the simulation.
func (r *ECS) Next() {
	// This would be a great place to introduce concurrency, but in order
	// to be able to compare this approach with others, all computations
	// are going to be performed linearly.
	for _, c := range r.chunks {
		for _, s := range c.Systems {
			for i := range c.Entities {
				s.Run(&c.Entities[i], &c.Entities)
			}
		}
	}
}

// Freeze exports the current state of the simulation.
func (r *ECS) Freeze() runner.State {
	// TODO: implement
	panic("")
}

// Restore sets the state of the simulation to one provided.
func (r *ECS) Restore(state runner.State, globals runner.Globals) {
	// TODO: implement
	panic("")
}
