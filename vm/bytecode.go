package vm

import(
	"fmt"
)

const (
	// Kinds
	StringKind symbol  = "string"
	NumberKind         = "number"
	BooleanKind        = "boolean"
	SymbolKind         = "symbol"
	FunctionKind       = "function"

	// Op codes
	Breakpoint         = "set-breakpoint"
	Set                = "set-register"
	Get                = "get-register"
	Save               = "save-register"
	Restore            = "restore-register"
	Label              = "label"
	Goto               = "goto"
	Test               = "test"
	Branch             = "branch"
)

type instruction struct{
	Op   symbol
	Args []interface{}
}

type symbol string

type value struct{
	kind symbol
	v    interface{}
}

func (val *value) String() string {
	return fmt.Sprintf("%s[%v]", val.kind, val.v)
}

func Instruction(op symbol, args ...interface{}) instruction {
	return instruction{op, args}
}

func Symbol(symbolName string) symbol {
	return symbol(symbolName)
}

func Value(v interface{}) *value {
	switch v.(type) {
	case int:
		return &value{NumberKind, v}
	case string:
		return &value{StringKind, v}
	case bool:
		return &value{BooleanKind, v}
	case *function:
		return &value{FunctionKind, v}
	default:
		panic(fmt.Sprintf("No support available for value type! %+v", v))
	}
}
