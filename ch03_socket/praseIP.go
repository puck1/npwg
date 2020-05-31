package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
	}
	dotAddr := os.Args[1]

	// net.ParseIP transfers a string (v4 or v6) to net.IP ([]byte).
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Fprintln(os.Stderr, "Invalid ip address")
		os.Exit(1)
	}

	// There is a function to create a mask from a 4-byte(hex) IPv4 address.
	//mask := net.IPv4Mask(0xff, 0xff, 0x00, 0x00)

	// DefaultMask method of a net.IP returns the default mask.
	mask := addr.DefaultMask()

	// A mask can be used by Mask method of a net.IP to find the net work for it.
	network := addr.Mask(mask)

	// Get number of mask bits and its leading ones.
	ones, bits := mask.Size()

	fmt.Fprintln(os.Stdout, "Address is ", addr.String(),
		" Mask length is ", bits,
		" Leading ones count is ", ones,
		" Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}
