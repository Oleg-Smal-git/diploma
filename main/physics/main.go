package main

import (
	"time"

	"github.com/Oleg-Smal-git/diploma/main/config"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
	"github.com/Oleg-Smal-git/diploma/services/physics/common"
)

func main() {
	// Initialize the components and set initial conditions.
	runner, archivist := initialize()
	state := common.State{
		Balls:             make([]*common.Ball, 0, config.StateCapacity),
		LastFrameDuration: time.Duration(0),
	}
	if err := archivist.LoadState(config.StateSource, &state); err != nil {
		panic("initialization failure: " + err.Error())
	}
	runner.Restore(&state, interfaces.Globals{
		FrameSimulationTime: config.FrameDuration,
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
