package vm

import(
	"fmt"
)

type interpreter struct{
	registers  map[symbol]*value
	labelTable map[symbol]int
}

type instruction struct{
	Op    symbol
	Args  []interface{}
}

func Instruction(op symbol, args ...interface{}) instruction {
	return instruction{op, args}
}

func Program(instructions ...instruction) []instruction {
	return instructions
}

type symbol string

func Symbol(symbolName string) symbol {
	return symbol(symbolName)
}

type value struct{
	kind symbol
	v    interface{}
}

func Value(v interface{}) *value {
	switch v.(type) {
	case int:
		return &value{Number, v}
	case string:
		return &value{String, v}
	case bool:
		return &value{Boolean, v}
	default:
		panic(fmt.Sprintf("No support available for value type! %+v", v))
	}
}

const (
	// Kinds
	String symbol      = "string"
	Number             = "number"
	Boolean            = "boolean"
	Func               = "function"

	// Special registers
	InstructionPointer = "instructionPointer"
	Accumulator        = "accumulator"

	// Op codes
	Breakpoint         = "set-breakpoint"
	Set                = "set-register"
	Get                = "get-register"
	Label              = "label"
	Goto               = "goto"
)

func newInterpreter() *interpreter {
	return &interpreter{make(map[symbol]*value), make(map[symbol]int)}
}

func (i *interpreter) setLabels(labelTable map[symbol]int) {
	i.labelTable = labelTable
}

func (i *interpreter) initialize() {
	// Reset interpreter state
	i.registers[InstructionPointer] = &value{Number, 0}
}

func (i *interpreter) position() int {
	return i.registers[InstructionPointer].v.(int)
}

func (i *interpreter) jump(pos int) {
	i.registers[InstructionPointer] = &value{Number, pos}
}

func (i *interpreter) advance() {
	i.registers[InstructionPointer] = &value{Number, i.position() + 1}
}

func (i *interpreter) do(inst instruction) {
	switch inst.Op {
	case Set:
		registerName := inst.Args[0].(symbol)
		registerValue := inst.Args[1].(*value)
		i.registers[registerName] = registerValue
		i.advance()
	case Get:
		registerName := inst.Args[0].(symbol)
		i.registers[Accumulator] = i.registers[registerName]
		i.advance()
	case Goto:
		destinationLabel := inst.Args[0].(symbol)
		if destinationPtr, exists := i.labelTable[destinationLabel]; exists {
			i.jump(destinationPtr)
		} else {
			panic(fmt.Sprintf("No such label `%s`", destinationLabel))
		}
	case Breakpoint:
		i.breakpoint()
	default:
		panic("Oh noes! What the hell kind of operation were you trying to perform?")
	}
}

func processLabels(labeledProgram []instruction) (map[symbol]int, []instruction) {
	var instructions []instruction
	labelTable := make(map[symbol]int)

	realI := 0
	for _, inst := range labeledProgram {
		if inst.Op == Label {
			// TODO: Check for duplicate labels
			labelTable[inst.Args[0].(symbol)] = realI
		} else {
			realI = realI + 1
			instructions = append(instructions, inst)
		}
	}

	return labelTable, instructions
}

func (v *VM) RunWithLabels(labeledProgram []instruction) {
	// TODO: Should this just be done by the compiler?
	labelTable, program := processLabels(labeledProgram)
	v.itpr.setLabels(labelTable)
	v.Run(program)
}

func (v *VM) Run(program []instruction) {
	v.itpr.initialize()
	for v.itpr.position() < len(program) {
		currentInstruction := program[v.itpr.position()]
		v.itpr.do(currentInstruction)
	}
}