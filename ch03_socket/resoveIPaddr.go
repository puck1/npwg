package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		// Tip: use `www.baidu.com` to test this program.
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	hostname := os.Args[1]

	/* IPAddr represents the address of an IP end point.
	type IPAddr struct {
		IP   IP
		Zone string // IPv6 scoped addressing zone
	}
	*/

	// net.ResolveIPAddr can perform DNS lookups on IP host names and returns a IPAddr,
	// Parameter "network" is one of "ip", "ip4" or "ip6".
	fmt.Println("[Start net.ResolveIPAddr]")
	addr, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	fmt.Println("[End net.ResolveIPAddr]")

	fmt.Println("[Start net.LookupHost]")
	// Hosts may have multiple IP addresses, so net.LookupHost can return a slice of IP STRING.
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Look up host error", err.Error())
		os.Exit(1)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
	fmt.Println("[End net.LookupHost]")

	fmt.Println("[Start net.LookupCNAME]")
	// If you wish to find the canonical name, use net.LookupCNAME.
	cname, err := net.LookupCNAME(hostname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Look up cname error", err.Error())
		os.Exit(1)
	}
	fmt.Println("CNAME is", cname)
	fmt.Println("[End net.LookupCNAME]")


	os.Exit(0)
}

