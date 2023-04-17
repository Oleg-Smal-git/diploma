package runner

import (
	"os"

	"github.com/vmihailenco/msgpack"
)

// LoadState sets the State from source file into target.
func LoadState(source string, target *State) error {
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(data, target)
}

// SaveState saves the source State in target file.
func SaveState(target string, source State) error {
	data, err := msgpack.Marshal(source)
	if err != nil {
		return err
	}
	return os.WriteFile(target, data, 0644)
}
