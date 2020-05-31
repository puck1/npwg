/* UDPDaytimeServer_v2
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

	// A UDP packet server.
	conn, err := net.ListenPacket("udp", service)
	checkError(err)

	for {
		// This is different from a TCP server
		handleClient(conn)
	}
}

func handleClient(conn net.PacketConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFrom(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteTo([]byte(daytime + "\n"), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1) }
}
