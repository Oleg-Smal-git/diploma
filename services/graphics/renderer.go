package graphics

import (
	"fmt"
	"os"
	"regexp"
	"sync"

	"github.com/Oleg-Smal-git/diploma/services/interfaces"

	"github.com/fogleman/gg"
)

// Confirm that Renderer satisfies interfaces.Renderer interface.
// This will throw a compile error otherwise.
var _ interfaces.Renderer = (*Renderer)(nil)

// Renderer implements interfaces.Renderer.
type Renderer struct {
	archivist      interfaces.Archivist
	context        *gg.Context
	workers        int
	files          chan string
	errors         chan error
	fileGroup      sync.WaitGroup
	errorGroup     sync.WaitGroup
	extensionRegex *regexp.Regexp
}

// NewRenderer instantiates a new Renderer.
func NewRenderer(archivist interfaces.Archivist, context *gg.Context, workers int) *Renderer {
	return &Renderer{
		archivist:      archivist,
		context:        context,
		workers:        workers,
		files:          nil,
		errors:         nil,
		fileGroup:      sync.WaitGroup{},
		errorGroup:     sync.WaitGroup{},
		extensionRegex: regexp.MustCompile(`(\.\w+)$`),
	}
}

// BulkRender renders all files in sourceDirectory and saves results to destinationDirectory.
func (r *Renderer) BulkRender(sourceDirectory string, destinationDirectory string, template interfaces.Renderable) error {
	entries, err := os.ReadDir(sourceDirectory)
	if err != nil {
		return err
	}
	r.files, r.errors = make(chan string), make(chan error)
	r.fileGroup.Add(r.workers)
	r.errorGroup.Add(1)
	go func() {
		defer r.errorGroup.Done()
		for e := range r.errors {
			r.consumeError(e)
		}
	}()
	for i := 0; i < r.workers; i++ {
		go func() {
			defer r.fileGroup.Done()
			for f := range r.files {
				if err := r.consumeInput(
					fmt.Sprintf("%v/%v", sourceDirectory, f),
					r.extensionRegex.ReplaceAllString(fmt.Sprintf("%v/%v", destinationDirectory, f), ".png"),
					template,
				); err != nil {
					r.errors <- err
				}
			}
		}()
	}
	for _, e := range entries {
		r.files <- e.Name()
	}
	close(r.files)
	r.fileGroup.Wait()
	close(r.errors)
	r.errorGroup.Wait()
	return nil
}

// Collect create an aggregation file (like .gif or .mp4).
func (r *Renderer) Collect(sourceDirectory string, destination string) error {
	// TODO: implement.
	panic("implement")
}

// consumeInput reads state file and renders the result.
func (r *Renderer) consumeInput(in string, out string, template interfaces.Renderable) error {
	object := template.New()
	if err := r.archivist.LoadState(in, &object); err != nil {
		return err
	}
	object.Render(r.context)
	return r.context.SavePNG(out)
}

// consumeError handles consumeInput errors.
func (r *Renderer) consumeError(err error) {
	fmt.Printf("Renderer.BulkRender failure: %s", err)
}
