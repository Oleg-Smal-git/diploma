package main

import (
	"github.com/Oleg-Smal-git/diploma/main/config"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

func main() {
	// Initialize the components and set initial conditions.
	runner, archivist := initialize()
	var (
		state interfaces.State
		err   error
	)
	if err = archivist.LoadState(config.StateSource, &state); err != nil {
		panic("initialization failure: " + err.Error())
	}
	runner.Restore(state, interfaces.Globals{
		FrameDuration: config.FrameDuration,
	})

	// Execute the simulation.
	for i := 0; i < config.FrameCap; i++ {
		runner.Next()
		runner.Freeze(&state)
		if err = archivist.SaveState(config.StateDestination, state); err != nil {
			panic("archivist failure: " + err.Error())
		}
	}
}
