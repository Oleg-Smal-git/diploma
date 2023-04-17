package archivist

import (
	"os"

	"github.com/Oleg-Smal-git/diploma/services/interfaces"

	"github.com/vmihailenco/msgpack"
)

// Confirm that Archivist satisfies interfaces.Archivist interface.
// This will throw a compile error otherwise.
var _ interfaces.Archivist = (*Archivist)(nil)

// Archivist is an interface used to interact with disk.
type Archivist struct{}

// NewArchivist instantiate an Archivist.
func NewArchivist() *Archivist {
	return &Archivist{}
}

// LoadState sets the State from source file into target.
func (Archivist) LoadState(source string, target *interfaces.State) error {
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(data, target)
}

// SaveState saves the source State in target file.
func (Archivist) SaveState(target string, source interfaces.State) error {
	data, err := msgpack.Marshal(source)
	if err != nil {
		return err
	}
	return os.WriteFile(target, data, 0644)
}
