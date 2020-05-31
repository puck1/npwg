/* getHeadInfo_v2
 */
package main
import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	// Instead of separate dial functions for TCP and UDP, you can use a single function Dial.
	// Note the parameter address is type STRING so that string resolving is not needed.
	// The Go net package recommends using these interface types rather than the concrete ones.
	// But by using them, you lose specific methods such as `SetKeepAlive` of TCPConn and `SetReadBuffer` of UDPConn,
	// unless you do a type cast. It is your choice.
	conn, err := net.Dial("tcp", service)
	checkError(err)

	defer conn.Close()

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1) }
}