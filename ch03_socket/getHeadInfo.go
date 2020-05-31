/* GetHeadInfo using TCP conn.
 */
package main
import (
	"net"
	"os"
	"fmt"
	//"io/ioutil"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	// Transfer string("host:port" or "ip:port") to a TCPAddr instance.
	// Parameter network is "tcp4", "tcp6" or "tcp".
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// Client uses DialTCP to create a connection.
	// Parameter laddr is local address and raddr is remote address.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// Write bytes to a conn.
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// Read results, you can also use conn.Read to read into a byte slice.
	// But note the connection may consist of several TCP packets,
	// so we need to keep reading till the end of file.
	// The io/ioutil function ReadAll will look after these issues and return the complete response.
	result := make([]byte, 512)
	_, err = conn.Read(result)
	//result, err := ioutil.ReadAll(conn)
	checkError(err)

	// Note only []byte can be coerced to string while [512]byte can't.
	fmt.Println(string(result))

	// Note error checking should be normal for network program.
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1) }
}