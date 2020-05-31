/* DaytimeServer
 */
package main
import (
	"fmt"
	//"fmt"
	"net"
	"os"

	//"os"
	"time"
)
func main() {
	service := ":1200"

	// Transfer string("host:port" or "ip:port") to a TCPAddr instance.
	// Parameter network is "tcp4", "tcp6" or "tcp".
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// The argument network can be set to one of the strings "tcp", "tcp4" or "tcp6".
	// The IP address should be set to zero if you want to listen on all network interfaces,
	// or to the IP address of a single network interface if you only want to listen on that interface.
	// If the port is set to zero, then the O/S will choose a port for you.
	listener, err := net.ListenTCP("tcp", tcpAddr)
	// If encounter any error in the initial stage, just shut down.
	// But once it starts listen, any error in interacting with clients won't shut down the server,
	// but to close the connection and continue to serve other clients.
	checkError(err)

	fmt.Printf("Server is listening at %s:%d... ", tcpAddr.IP.String(), tcpAddr.Port)
	// Run server forever
	for {
		// Accept will block the server until get a request and return a Conn.
		conn, err := listener.Accept()
		// The server should run forever, so that if any error occurs with a client,
		// the server just ignores that client and carries on.
		// A client could otherwise try to mess up the connection with the server, and bring the conn down!
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime + "\n"))	// don't care about return value conn.Close()
		conn.Close()						// we're finished with this client
	}
}

// Just open up a telnet connection to that host:
// telnet localhost 1200

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1) }
}