package graphics

import (
	"github.com/Oleg-Smal-git/diploma/services/interfaces"
)

// Confirm that Renderer satisfies interfaces.Renderer interface.
// This will throw a compile error otherwise.
var _ interfaces.Renderer = (*Renderer)(nil)

// Renderer implements interfaces.Renderer.
type Renderer struct{}

// NewRenderer instantiates a new Renderer.
func NewRenderer() *Renderer {
	return &Renderer{}
}

// BulkRender renders all files in sourceDirectory and saves results to destinationDirectory.
func (r *Renderer) BulkRender(sourceDirectory string, destinationDirectory string) {
	//TODO implement me
	panic("implement me")
}

// Collect create an aggregation file (like .gif or .mp4).
func (r *Renderer) Collect(sourceDirectory string, destination string) {
	//TODO implement me
	panic("implement me")
}
