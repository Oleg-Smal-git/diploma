package main

import (
	"fmt"
	"time"

	"github.com/Oleg-Smal-git/diploma/main/config"
	"github.com/Oleg-Smal-git/diploma/services/archivist"
	"github.com/Oleg-Smal-git/diploma/services/graphics"
	"github.com/Oleg-Smal-git/diploma/services/instances"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

func initialize() interfaces.Renderer {
	return graphics.NewRenderer(archivist.NewArchivist(), config.ImageWidth, config.ImageHeight, config.FramesPerSecond, config.GraphicsWorkerPool)
}

func main() {
	// Initialize renderer.
	renderer := initialize()
	// Render individual frames.
	if err := renderer.BulkRender(config.StateDestination, config.FrameDestination, &instances.State{}); err != nil {
		panic(err)
	}
	// Collect frames into aggregation.
	if err := renderer.Collect(config.FrameDestination, fmt.Sprintf("%v/%v.avi", config.AggregationDestination, time.Now().UnixNano())); err != nil {
		panic(err)
	}
}
