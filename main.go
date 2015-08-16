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

	instance.Initialize()

	// prog := vm.Program(
	// 	// Initialize `a` = 10
	// 	vm.Instruction(vm.Set, vm.Symbol("a"), vm.Value(10)),
	// 	// vm.Instruction(vm.Breakpoint, vm.Value("program.r @ line 1")),

	// 	// Set looping point `loop1` here
	// 	vm.Instruction(vm.Label, vm.Symbol("loop1")),
	// 	// Update `a` = 20
	// 	vm.Instruction(vm.Set, vm.Symbol("a"), vm.Value(20)),
	// 	// Go back to loop point `loop1`
	// 	vm.Instruction(vm.Goto, vm.Symbol("loop1")),
	// )

	instance.Load(vm.Symbol("_main"), vm.Function(
		// Set a = 10
		vm.Instruction(vm.Set, vm.Symbol("a"), vm.Value(10)),
		vm.Instruction(vm.Breakpoint, vm.Value("program.r @ line 1")),

		// Update a = 20
		vm.Instruction(vm.Set, vm.Symbol("a"), vm.Value(20)),

		// Update a = 30, set b = -10
		vm.Instruction(vm.Set, vm.Symbol("a"), vm.Value(30)),
		vm.Instruction(vm.Set, vm.Symbol("b"), vm.Value(-10)),
	))

	instance.RunWithLabels(entryPoint)
}
