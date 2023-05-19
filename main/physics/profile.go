//go:build BUILD_PROFILE

package main

import (
	"os"

	"github.com/Oleg-Smal-git/diploma/config"
	"github.com/Oleg-Smal-git/diploma/services/instances"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"

	"runtime/pprof"
)

func solve(runner interfaces.Runner, archivist interfaces.Archivist, state *instances.State) {
	var (
		err         error
		memory, cpu *os.File
	)
	memory, err = os.Create(config.MemoryProfileDestination)
	cpu, err = os.Create(config.CPUProfileDestination)
	defer memory.Close()
	defer cpu.Close()
	if err = pprof.StartCPUProfile(cpu); err != nil {
		panic(err)
	}
	for i := 0; i < config.FrameCap; i++ {
		runner.Next()
	}
	pprof.StopCPUProfile()
	if err = pprof.WriteHeapProfile(memory); err != nil {
		panic(err)
	}
}
