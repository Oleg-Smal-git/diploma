package main

import (
	"github.com/Oleg-Smal-git/diploma/services/graphics"
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

func initialize() interfaces.Renderer {
	return graphics.NewRenderer()
}

func main() {
	_ = initialize()
}
