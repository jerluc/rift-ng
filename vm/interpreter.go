package vm

import(
	"fmt"
)

type machine struct{
	registers  map[symbol]*value
}

func newMachine() *machine {
	return &machine{make(map[symbol]*value)}
}

func (m *machine) initialize() {
	// Reset machine state
	m.setRegister(InstructionPointer, Value(0))
}

func (m *machine) position() int {
	return m.getRegister(InstructionPointer).v.(int)
}

func (m *machine) jump(pos int) {
	m.setRegister(InstructionPointer, Value(pos))
}

func (m *machine) advance() {
	m.setRegister(InstructionPointer, Value(m.position() + 1))
}

func (m *machine) do(inst instruction) {
	switch inst.Op {
	case Set:
		registerName := inst.Args[0].(symbol)
		registerValue := inst.Args[1].(*value)
		m.setRegister(registerName, registerValue)
		m.advance()
	case Get:
		registerName := inst.Args[0].(symbol)
		m.setRegister(Accumulator, m.getRegister(registerName))
		m.advance()
	case Goto:
		destinationLabel := inst.Args[0].(symbol)
		if destinationPtr, exists := m.labelTable[destinationLabel]; exists {
			m.jump(destinationPtr)
		} else {
			panic(fmt.Sprintf("No such label `%s`", destinationLabel))
		}
	// case Breakpoint:
	// 	m.breakpoint(inst)
	default:
		panic("Oh noes! What the hell kind of operation were you trying to perform?")
	}
}

func (v *VM) Load(name symbol, f *function) {
	m.setRegister(name, Value(f))
}

func (v *VM) Initialize() {
	v.itpr.initialize()
}

func (v *VM) Run(entryPoint symbol) {
	f := v.itpr.getRegister(entryPoint).v.(*function)
	for v.itpr.position() < f.Len() {
		v.itpr.do(f.Inst(v.itpr.position()))
	}
}