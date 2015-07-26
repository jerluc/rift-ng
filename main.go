package main

import (
	"flag"
	"github.com/jerluc/rift/vm"
)

func main() {
	flag.Parse()
	instance := vm.NewVM()
	for _, rcFile := range(flag.Args()) {
		instance.LoadFile(rcFile)
	}
	instance.Listen()
}
