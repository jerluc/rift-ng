package vm

const (
	// Special registers
	InstructionPointer symbol = "_instructionPointer"
	Accumulator               = "_accumulator"
)

func (i *interpreter) setRegister(registerName symbol, registerValue *value) {
	i.registers[registerName] = registerValue
}

func (i *interpreter) getRegister(registerName symbol) *value {
	return i.registers[registerName]
}

func (i *interpreter) getRegisterNames() []string {
	var registerNames []string
	for registerName, _ := range i.registers {
		registerNames = append(registerNames, string(registerName))
	}
	return registerNames
}
