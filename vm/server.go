package vm

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Binds the VM server to a TCP socket
// and starts listening for incoming VM
// requests for both local and remote
// RPC operations.
//
// TODO: Is there any benefit to using a
// non-TCP socket (UDP/QUIC/etc)?
//
// TODO: Allow port to be configurable or
// possibly even ephemeral (assuming some
// kind of service discovery could be
// used in the latter case)
func (v *VM) Listen() {
	ln, _ := net.Listen("tcp", ":11117")
	defer ln.Close()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down")
		v.dumpState()
		done <- true
		ln.Close()
	}()

	for {
		select {
		case <-done:
			return
		default:
			req, _ := ln.Accept()
			go v.HandleRequest(req)
		}
	}
}

// Handles an incoming VM server request
//
// TODO: Actually establish some kind of
// light-weight protocol for handling
// multiple kinds of VM RCP operations
func (v *VM) HandleRequest(req net.Conn) {
	defer req.Close()
	v.Load(bufio.NewReader(req))
}
