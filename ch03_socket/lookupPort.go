/* LookupPort
 */
package main
import (
	"net"
	"os"
	"fmt"
)
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type service\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	// On a Unix system, the commonly used ports are listed in the file /etc/services.
	// Go has a function to interrogate this file.
	// net.LookupPort returns the port the service uses on specified network type.
	// The parameter networkType is "tcp" or "udp", service is "telnet" or "domain"(for DNS) and so on.
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	fmt.Println("Service port ", port)
	os.Exit(0)
}
