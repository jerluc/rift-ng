package vm

import (
	"fmt"
)

const (
	Advance  = "advance"
	Jump     = "jump"
	Show     = "show"
	ShowAll  = "showall"
	Continue = "continue"
)

func (i *interpreter) dumpRegisterContents() {
	fmt.Println("Registers: {")
	for k, v := range i.registers {
		fmt.Printf("\t%s => %+v\n", k, v)
	}
	fmt.Println("}")
}

func (i *interpreter) breakpoint() {
	fmt.Println("Hit breakpoint @", i.position())
	for {
		var cmd string
		fmt.Scanf("%s", &cmd)
		switch cmd {
		case Advance:
			// TODO: Actually interpret but stay in here!
		case Jump:
			// TODO: Actually go to the provided label
		case Show:
			// TODO: Actually show the contents of one register
		case ShowAll:
			i.dumpRegisterContents()
		case Continue:
			i.advance()
			return
		default:
			fmt.Printf("No such command \"%s\".\n", cmd)
			fmt.Println("Must be one of:", Advance, Jump, Show, ShowAll, Continue)
		}
	}
}