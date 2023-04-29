package ecs

import (
	"time"

	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

// Confirm that ECS satisfies interfaces.Runner interface.
// This will throw a compile error otherwise.
var _ interfaces.Runner = (*ECS)(nil)

// ECS stands for Entity Component System and is an architectural
// pattern we will be using for this implementation of interfaces.interfaces.
type ECS struct {
	globals interfaces.Globals
	chunks 			   []Chunk
	lastFrameStartTime time.time
	lastFrameEndTime   time.Time
	lastFrameDuration  time.Duration
}

// NewRunner constructs an ECS object.
func NewRunner(componentRegistrar []Component, archetypesRegistrar []ComponentID, systemRegistrar []System, globals interfaces.Globals) *ECS {
	ecs := ECS{
		globals: globals,
		chunks:  make([]Chunk, 0, len(archetypesRegistrar)),
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
				chunk.Systems[len(chunk.Systems)-1].Restore(&ecs.globals)
			}
		}
		ecs.chunks[a] = chunk
	}
	return &ecs
}

// Next performs one atomic step of the simulation.
func (r *ECS) Next() {
	r.lastFrameStartTime = time.Now()
	// This would be a great place to introduce concurrency, but in order
	// to be able to compare this approach with others, all computations
	// are going to be performed linearly.
	for _, c := range r.chunks {
		for _, s := range c.Systems {
			for i := range c.Entities {
				s.Run(&i, &c.Entities[i], &c.Entities)
			}
		}
	}
	r.lastFrameEndTime = time.Now()
	r.lastFrameDuration = r.lastFrameEndTime.Sub(r.lastFrameStartTime)
}

// Freeze exports the current state of the simulation.
func (r *ECS) Freeze(state *interfaces.State) {
	state.LastFrameDuration = r.lastFrameDuration
}

// Restore sets the state of the simulation to one provided.
func (r *ECS) Restore(state interfaces.State, globals interfaces.Globals) {
	// TODO: implement
	panic("")
}
