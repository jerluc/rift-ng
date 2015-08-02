package vm

// import(
// 	"container/list"
// 	"fmt"
// 	"strconv"
// )

// type interpreter struct{
// 	stack *list.List
// 	env   map[string]*value
// }

// type Instruction struct{
// 	Op    string
// 	Args  []interface{}
// }

// type value struct{
// 	kind string
// 	v    interface{}
// }

// const (
// 	NEW  = "new"
// 	LET  = "let"
// 	REF  = "ref"
// 	CALL = "call"
// 	JUMP = "jump"
// )

// func newInterpreter() *interpreter {
// 	return &interpreter{list.New(), make(map[string]*value)}
// }

// func (i *interpreter) push(v *value) {
// 	fmt.Printf("Pushing value %+v\n", v)
// 	i.stack.PushFront(v)
// }

// func (i *interpreter) pop() *value {
// 	if i.stack.Len() == 0 {
// 		panic("Empty stack!")
// 	}

// 	v := i.stack.Remove(i.stack.Front()).(*value)
// 	fmt.Printf("Popping value %+v\n", v)
// 	return v
// }

// func (i *interpreter) construct(valueKind string, ctorArgs []interface{}) {
// 	switch valueKind {
// 	default:
// 		panic(fmt.Sprintf("Invalid value kind [%s]", valueKind))
// 	case "n":
// 		num, _ := strconv.Atoi(ctorArgs[0].(string))
// 		i.push(&value{"n", num})
// 	case "s":
// 		i.push(&value{"s", ctorArgs[0].(string)})
// 	case "b":
// 		boolean, _ := strconv.ParseBool(ctorArgs[0].(string))
// 		i.push(&value{"b", boolean})
// 	}
// }

// // TODO: Should this bind and push the value back
// // onto the stack instead?
// func (i *interpreter) bind(refName string) {
// 	if _, exists := i.env[refName]; exists == true {
// 		panic("Ref already exists! " + refName)
// 	}

// 	i.env[refName] = i.pop()
// }

// func (i *interpreter) dereference(refName string) {
// 	if _, exists := i.env[refName]; exists != true {
// 		panic("Ref does not exist! " + refName)
// 	}

// 	i.push(i.env[refName])
// }

// func (i *interpreter) call(refName string) {
// 	// TODO: Take from the environment instead!

// 	switch refName {
// 	case "addn":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"n", left + right})
// 	case "adds":
// 		left := i.pop().v.(string)
// 		right := i.pop().v.(string)
// 		i.push(&value{"s", left + right})
// 	case "subn":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"n", left - right})
// 	case "muln":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"n", left * right})
// 	case "divn":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"n", left / right})
// 	case "modn":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"n", left % right})
// 	case "ltn":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"b", left < right})
// 	case "lten":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"b", left <= right})
// 	case "gtn":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"b", left > right})
// 	case "gten":
// 		left := i.pop().v.(int)
// 		right := i.pop().v.(int)
// 		i.push(&value{"b", left >= right})
// 	case "eq":
// 		left := i.pop().v
// 		right := i.pop().v
// 		i.push(&value{"b", left == right})
// 	case "neq":
// 		left := i.pop().v
// 		right := i.pop().v
// 		i.push(&value{"b", left != right})
// 	}
// }

// func (v *VM) Start(instructions ...Instruction) {
// 	for _, i := range instructions {
// 		switch i.Op {
// 		case NEW:
// 			v.itpr.construct(i.Args[0].(string), i.Args[1:])
// 		case LET:
// 			v.itpr.bind(i.Args[0].(string))
// 		case REF:
// 			v.itpr.dereference(i.Args[0].(string))
// 		case CALL:
// 			v.itpr.call(i.Args[0].(string))
// 		// case JUMP:
// 		// 	v.itpr.jump(i.Args[0].(string))	
// 		}
// 	}

// 	fmt.Println("Final interpreter environment:")
// 	for k, v := range v.itpr.env {
// 		fmt.Printf("\t%s => %+v\n", k, v)
// 	}
// }