package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

// Whatever type it is...
type Arith int

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// Note `reply` must be a pointer.
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
	s := rpc.NewServer()
	err := s.Register(arith)						// Note Register receive a pointer to an object.
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":1200", s)		// rpc.Server implements http.Handler interface.
	if err != nil {
		log.Fatal(err)
	}
	// To use rpc.DefaultServer:
	//rpc.Register(arith)	// Register rpc functions to rpc.DefaultServer.
	//rpc.HandleHTTP()		// Must be called when use rpc.DefaultServer and http.DefaultMux
							// If use http.ListenAndServe(":1200", rpc.DefaultServer) then HandlerHTTP is not necessary.
	//http.ListenAndServe(":1200", nil)		// Use http.DefaultMux as handler/mux.
}
