package main

import (
	"github.com/jerluc/rift/vm"
)

func main() {
	instance := vm.NewVM()
	instance.Load("hello.rc")
}
