package vm

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// 	"github.com/bobappleyard/readline"
// )

// const (
// 	Step          = "step"
// 	Jump          = "jump"
// 	GetRegister   = "get"
// 	SetRegister   = "set"
// 	DumpRegisters = "dump"
// 	Continue      = "continue"
// 	// TODO: Implement "clear-breakpoints"
// )

// var DebuggerCommands = []string{Step, Jump, GetRegister, SetRegister, DumpRegisters, Continue}

// func (i *interpreter) dumpRegisterContents() {
// 	fmt.Println("Registers: {")
// 	for k, v := range i.registers {
// 		fmt.Printf("\t%s => %+v\n", k, v)
// 	}
// 	fmt.Println("}")
// }

// func matchesPrefix(options []string, prefix string) []string {
// 	var possibilities []string
// 	for _, opt := range options {
// 		if strings.HasPrefix(opt, prefix) {
// 			possibilities = append(possibilities, opt)
// 		}
// 	}
// 	return possibilities
// }

// func (i *interpreter) breakpoint(b instruction) {
// 	fmt.Printf("Hit breakpoint %+v\n", b.Args)
// 	readline.Completer = func(query, ctx string) []string {
// 		if query == ctx {
// 			return matchesPrefix(DebuggerCommands, query)
// 		} else {
// 			cmdParts := strings.Fields(ctx)
// 			switch cmdParts[0] {
// 			case SetRegister:
// 				fallthrough
// 			case GetRegister:
// 				if len(cmdParts) == 1 {
// 					return i.getRegisterNames()	
// 				}
// 				return matchesPrefix(i.getRegisterNames(), cmdParts[1])
// 			}
// 		}
// 		return []string{}
// 	}
// 	i.advance()
// 	for {
// 		cmd, err := readline.String("riftdb> ")
// 		if err == io.EOF {
// 			break
// 		}
// 		readline.AddHistory(cmd)

// 		cmdParts := strings.Fields(cmd)

// 		switch cmdParts[0] {
// 		case Step:
// 			if i.running() {
// 				i.process()
// 			} else {
// 				return
// 			}
// 		case Continue:
// 			return
// 		case Jump:
// 			// TODO: Actually go to the provided label
// 		case GetRegister:
// 			if len(cmdParts) == 2 {
// 				val := i.getRegister(symbol(cmdParts[1]))
// 				fmt.Printf("%s = %v\n", cmdParts[1], val)
// 			} else {
// 				fmt.Println("Usage: get [register-name]")
// 			}
// 		case SetRegister:
// 			if len(cmdParts) == 3 {
// 				i.setRegister(symbol(cmdParts[1]), Value(cmdParts[2]))
// 			} else {
// 				fmt.Println("Usage: set [register-name] [register-value]")
// 			}
// 		case DumpRegisters:
// 			i.dumpRegisterContents()
// 		case "":
// 		default:
// 			fmt.Printf("No such command \"%s\".\n", cmd)
// 			fmt.Println("Must be one of:", DebuggerCommands)
// 		}
// 	}
// }