package main

import (
	"github.com/Oleg-Smal-git/diploma/main/config"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
	"time"
)

func main() {
	// Initialize the components and set initial conditions.
	runner, archivist := initialize()
	state := interfaces.State{
		Balls:         make([]interfaces.Ball, 0, config.StateCapacity),
		LastFrameTime: time.Duration(0),
	}
	if err := archivist.LoadState(config.StateSource, &state); err != nil {
		panic("initialization failure: " + err.Error())
	}
	runner.Restore(state, interfaces.Globals{
		FrameDuration: config.FrameDuration,
	})

	// Execute the simulation.
	for i := 0; i < config.FrameCap; i++ {
		runner.Next()
		runner.Freeze(&state)
		if err := archivist.SaveState(config.StateDestination, state); err != nil {
			panic("archivist failure: " + err.Error())
		}
	}
}
