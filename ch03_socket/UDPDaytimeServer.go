/* UDPDaytimeServer
 */
package main
import (
	"fmt"
	"net"
	"os"
	"time"
)
func main() {
	service := ":1200"

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	// net.ListenUDP will NOT block server though it returns a connection type.
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		// This is different from a TCP server.
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	// conn.ReadFromUDP WILL block server.
	// ReadFromUDP will return client addr.
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	// WriteToUDP need client addr to send response.
	conn.WriteToUDP([]byte(daytime + "\n"), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1) }
}
