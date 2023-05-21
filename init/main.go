package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"

	"github.com/Oleg-Smal-git/diploma/config"
	"github.com/Oleg-Smal-git/diploma/services/archivist"
	"github.com/Oleg-Smal-git/diploma/services/instances"
)

func main() {
	overrideN := flag.Int("overrideN", 0, "")
	flag.Parse()
	var n int
	if *overrideN != 0 {
		n = *overrideN
	} else {
		n = defaultN
	}
	arch := archivist.NewArchivist(config.MarshalFunctor, config.UnmarshalFunctor)
	state := instances.State{
		Balls:             make([]*instances.Ball, 0),
		LastFrameDuration: 0,
	}
	r := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 0; i < n; i++ {
		col, row := i%r, i/r
		state.Balls = append(state.Balls, &instances.Ball{
			X:      float64(xMin + (xMax-xMin)*(col+1)/(r+1)),
			Y:      float64(yMin + (yMax-yMin)*(row+1)/(r+1)),
			Radius: radius,
			SpeedX: 2 * speed * (rand.Float64() - 0.5),
			SpeedY: 2 * speed * (rand.Float64() - 0.5),
		})
	}
	if err := arch.SaveState(fmt.Sprintf("./buff/start"), state); err != nil {
		panic(err)
	}
}
