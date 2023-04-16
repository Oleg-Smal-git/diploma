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
		SpeedX, SpeedY float64
	}
)
