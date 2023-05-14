package config

import (
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

const (
	// StateSource is a path to initial state source file.
	StateSource = "./buff/start.mpk"
	// StateDestination is a directory to which intermediate states are saved.
	StateDestination = "./buff/mpk"
	// FrameDestination is a directory to which individual frames are saved.
	FrameDestination = "./buff/png"
	// AggregationDestination is a directory to which the end results are saved.
	AggregationDestination = "./buff/gif"
	// StateCapacity describes the max amount of entities to be stored in state.
	// Used to pre-allocate memory during initialization.
	StateCapacity = 100
	// FrameCap is the amount of frames after which the simulation stops.
	FrameCap = 3600
	// FrameDuration is the amount of imaginary time that a frame lasts.
	// for simplicity's sake, all the numbers here are calibrated around
	// this variable being evaluated in seconds.
	FrameDuration = 1. / 60
	// ImageWidth describes frame size.
	ImageWidth = 1024
	// ImageHeight describes frame size.
	ImageHeight = 1024
	// GraphicsWorkerPool is the size of the concurrent worker bucket for renderer.
	GraphicsWorkerPool = 10
)

var (
	Globals = interfaces.Globals{
		FrameSimulationTime: FrameDuration,
		Boundaries: struct{ MinX, MaxX, MinY, MaxY float64 }{
			MinX: 0,
			MaxX: ImageWidth,
			MinY: 0,
			MaxY: ImageHeight,
		},
	}
)
