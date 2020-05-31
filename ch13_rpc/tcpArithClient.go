package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: " + os.Args[0] + "host:port\n")
		os.Exit(1)
	}

	// The only difference from http is that use rpc.Dial instead of rpc.DialHTTP
	client, err := rpc.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	args := Args{17, 8}

	var reply int
	// Synchronous call.
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d/%d=%d, remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}

