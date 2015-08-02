package vm

import (
	"fmt"
	"os"
	"io"
	"github.com/jerluc/rift/rc"
)

type VM struct{
	state *vmState
	itpr  *interpreter
}

// Constructs a new VM instance, doing all the
// usual initialization stuff
func NewVM() *VM {
	return &VM{newVMState(), newInterpreter()}
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

	v.state.load(rDef)

	fmt.Printf("Loaded RDef [%s]\n", rDef.Name)
}

func (v *VM) dumpState() {
	fmt.Printf("VM RTable\n---------\n")
	for name, rDef := range v.state.copy() {
		fmt.Printf("\t%s:", name)
		for i := 0; i < len(rDef.Code); i++ {
			if i % 256 == 0 {
				fmt.Print("\n\t")
			}
			fmt.Printf("%c", rDef.Code[i])
		}
		fmt.Println("\n")
	}
}
