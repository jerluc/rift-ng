package vm

import (
	"sync"
	"github.com/jerluc/rift/rc"
)

// TODO: Interesting... https://golang.org/src/runtime/malloc.go?s=16145:16202#L533

type vmState struct{
	rtable map[string]*rc.RDef
	m      *sync.RWMutex
}

func newVMState() *vmState {
	return &vmState{
		rtable: make(map[string]*rc.RDef),
		m:      &sync.RWMutex{},
	}
}

func (v *vmState) retrieve(name string) (*rc.RDef, bool) {
	defer v.m.RUnlock()
	v.m.RLock()
	rdef, found := v.rtable[name]
	return rdef, found
}

func (v *vmState) load(rdef *rc.RDef) {
	defer v.m.Unlock()
	v.m.Lock()
	v.rtable[rdef.Name] = rdef
}

func (v *vmState) copy() map[string]*rc.RDef {
	defer v.m.RUnlock()
	v.m.RLock()
	clone := make(map[string]*rc.RDef)
	for k, v := range v.rtable {
		clone[k] = v
	}

	return clone
}

