package vm

import (
	"net"
)

func (v *VM) Listen() {
	ln, _ := net.Listen("tcp", ":11117")
	for {
		conn, _ := ln.Accept()
		go v.Load(conn)
	}
}
