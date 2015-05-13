package vm

import (
	"fmt"
	"github.com/jerluc/rift/rc"
)

type VM struct{
	env map[string]interface{}
}

// Constructs a new VM instance, doing all the
// usual initialization stuff
func NewVM() *VM {
	return &VM{make(map[string]interface{})}
}

// Loads an RC file at the given file path, and
// then loads the RC definitions into the current
// VM instance.
//
// Note that this should atomically replace all
// active definitions of presently-referenced
// dispatch points.
func (v *VM) Load(rcFilePath string) {
	rcFile, loadErr := rc.LoadFile(rcFilePath)
	if loadErr != nil {
		// TODO: Obviously do something better
		panic("Ah shit")
	}

	fmt.Printf("I'm an RC file! %+v\n", rcFile)
}