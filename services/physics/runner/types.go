package runner

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
		Freeze() State
		// Restore sets the State and Globals of the simulation to one provided.
		Restore(State, Globals)
	}

	// State represents an exhaustive description of physics state of the simulation.
	State struct {
		// Balls is a collection of all Ball objects that take part in the simulation.
		Balls []Ball
		// LastFrameTime is the amount of time it took to compute last frame.
		LastFrameTime time.Duration
	}

	// Ball goes bounce :)
	Ball struct {
		X, Y           float64
		Radius         float64
		SpeedX, SpeedY float64
	}

	// Globals is a wrapper for all simulation config values, like frame duration.
	Globals struct {
		FrameDuration float64
	}
)
