package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// In go rpc package, name of structure can be different from the server, eg.Args or Values,
// but the field name must stay the same, if field names are `C,B int`, then C will be ignored and A will be 0.
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: " + os.Args[0] + " host:port\n")
		os.Exit(1)
	}

	client, err := rpc.DialHTTP("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	args := Args{17, 8}

	var reply int
	endAsync := make(chan interface{}, 1)		// Block program until asynchronous call finishes.
	// Asynchronous call.
	// It seems no influence that you pass an object or a pointer of `args`, but `reply` must be a pointer.
	// Note the method name must be completed.
	divCall := client.Go("Arith.Multiply", args, &reply, nil)
	go func(){
		repCall := <-divCall.Done
		// Because Reply is type interface{} and is a pointer to reply, so use type assert.
		reply = *repCall.Reply.(*int)
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
		endAsync <- struct {}{}
	}()

	var quo Quotient
	// Synchronous call.
	err = client.Call("Arith.Divide", &args, &quo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d/%d=%d, remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)

	<-endAsync			// Wait.
}
