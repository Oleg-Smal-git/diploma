package main

import (
	"errors"
	"flag"
	"os"
	"time"

	"github.com/Oleg-Smal-git/diploma/config"
	"github.com/Oleg-Smal-git/diploma/services/instances"
)

func main() {
	// Initialize the components and set initial conditions.
	runner, archivist := initialize()
	state := instances.State{
		Balls:             make([]*instances.Ball, 0, config.StateCapacity),
		LastFrameDuration: time.Duration(0),
	}
	source := flag.String("state_source", "", "")
	flag.Parse()
	if *source == "" {
		*source = config.StateSource
	}
	if err := archivist.LoadState(*source, &state); err != nil {
		panic("initialization failure: " + err.Error())
	}
	runner.Restore(&state, &config.Globals)
	// Copy initial state as first result.
	if _, err := os.Stat(config.StateDestination); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(config.StateDestination, os.ModePerm)
		if err != nil {
			panic("initialization failure: " + err.Error())
		}
	}
	// Execute the simulation.
	solve(runner, archivist, &state)
}
