package vm

import (
	"fmt"
	"io"
	"os"
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
func (v *VM) LoadFile(rcFilePath string) {
	file, openErr := os.Open(rcFilePath)
	if openErr != nil {
		panic(openErr)
	}

	v.Load(file)
}

func (v *VM) Load(rcStream io.Reader) {
	rDef, loadErr := rc.LoadStream(rcStream)
	if loadErr != nil {
		// TODO: Obviously do something better
		panic(loadErr)
	}

	fmt.Printf("Loaded RDef %+v\n", rDef)
}
