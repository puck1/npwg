package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// jsonrpc just use json instead of gob. As such, client or servers could be
// written in other language that understand sockets and JSON.

type Arith int

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func (t *Arith) Multiply (args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide (args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)

	// Create a rpc server and register.
	s := rpc.NewServer()
	err := s.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Resolve tcp address and create listener.
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1200")
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			continue
		}
		s.ServeCodec(jsonrpc.NewServerCodec(conn))
		// If use rpc.DefaultServer: jsonrpc.ServeConn(conn)
	}
}
