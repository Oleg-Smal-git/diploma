package config

const (
	// StateSource is a path to initial state source file.
	StateSource = "./state.mpk"
	// StateDestination is a path to which intermediate states are saved.
	StateDestination = ""
	// FrameCap is the amount of frames after which the simulation stops.
	FrameCap = 3600
	// FrameDuration is the amount of imaginary time that a frame lasts.
	// for simplicity's sake, all the numbers here are calibrated around
	// this variable being evaluated in seconds.
	FrameDuration = 1. / 60
)
