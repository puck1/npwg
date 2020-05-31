package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

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

	// Create rpc server and register
	s := rpc.NewServer()
	err := s.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Resolve tcp address and create listener
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1200")
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// rpc.Accept of rpcServer.Accept encapsulates listener.Accept WITH LOOP so we don't need to use for-loop here.
	//s.Accept(l)
	// If use rpc.DefaultServer: rpc.Accept(l)
	// same:
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			continue
		}
		s.ServeConn(conn)
		// If use default rpc server: rpc.ServeConn(conn)
	}
}
