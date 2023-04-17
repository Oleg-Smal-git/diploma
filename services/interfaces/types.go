package interfaces

import (
	"time"
)

type (
	// Runner is an interface that represents an object that performs all
	// the physics calculations, regardless of how they are implemented.
	Runner interface {
		// Next performs one atomic step of the simulation.
		Next()
		// Freeze exports the current state of the simulation.
		Freeze(*State)
		// Restore sets the State and Globals of the simulation to one provided.
		Restore(State, Globals)
	}

	// Archivist is an interface used to interact with disk.
	Archivist interface {
		// LoadState sets the State from source file into target.
		LoadState(source string, target *State) error
		// SaveState saves the source State in target file.
		SaveState(target string, source State) error
	}

	// State represents an exhaustive description of physics state of the simulation.
	State struct {
		// Balls is a collection of all Ball objects that take part in the simulation.
		Balls []Ball
		// LastFrameTime is the amount of time it took to compute last frame.
		LastFrameTime time.Duration
	}

	// Globals is a wrapper for all simulation config values, like frame duration.
	Globals struct {
		FrameDuration float64
	}

	// Ball goes bounce :)
	Ball struct {
		X, Y           float64
		Radius         float64
		SpeedX, SpeedY float64
	}
)
