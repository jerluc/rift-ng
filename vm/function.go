package vm

type function struct{
	instructions []instruction
	labelTable   map[symbol]int
	breakpoints  map[int]*value
}

// Given a set of instructions for a function,
// does a preprocess pass through to build out
// the function's local label table and breakpoint
// table (provided that debug mode is enabled).
func processLabelsAndBreakpoints(originalInstructions []instruction, debugEnabled bool) ([]instruction, map[symbol]int, map[int]*value) {
	var instructions []instruction
	labelTable := make(map[symbol]int)
	breakpoints := make(map[int]*value)

	realI := 0
	for _, inst := range originalInstructions {
		switch inst.Op {
		case Label:
			// TODO: Check for duplicate labels
			labelTable[inst.Args[0].(symbol)] = realI
		case Breakpoint:
			if debugEnabled {
				breakpoints[realI] = inst.Args[0].(*value)
			}
		default:
			realI = realI + 1
			instructions = append(instructions, inst)
		}
	}

	return instructions, labelTable, breakpoints
}

func Function(instructions ...instruction) *function {
	// TODO: Maybe move `debugEnabled = true` elsewhere?
	processedInstructions, labelTable, breakpoints := processLabelsAndBreakpoints(instructions, true)
	return &function{processedInstructions, labelTable, breakpoints}
}

func (f *function) Inst(pos int) instruction {
	return f.instructions[pos]
}

func (f *function) Len() int {
	return len(f.instructions)
}