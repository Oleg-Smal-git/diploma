//go:build BUILD_EXPORT

package main

import (
	"fmt"
	"github.com/Oleg-Smal-git/diploma/config"
	"github.com/Oleg-Smal-git/diploma/services/instances"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

func solve(runner interfaces.Runner, archivist interfaces.Archivist, state *instances.State) {
	if err := archivist.SaveState(fmt.Sprintf("%v/%v", config.StateDestination, 0), state); err != nil {
		panic("archivist failure: " + err.Error())
	}
	for i := 0; i < config.FrameCap; i++ {
		runner.Next()
		runner.Freeze(state)
		if err := archivist.SaveState(fmt.Sprintf("%v/%v", config.StateDestination, i+1), *state); err != nil {
			panic("archivist failure: " + err.Error())
		}
	}
}
