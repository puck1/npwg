package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	rootCAFile := "../X509Cert/localhost.pem"
	rootCABytes, err := ioutil.ReadFile(rootCAFile)
	checkError(err)

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(rootCABytes) {
		log.Fatal("Cannot get rootCA")
	}

	config := tls.Config{
		// InsecureSkipVerify用来控制客户端是否证书和服务器主机名。
		// 如果设置为true,则不会校验证书以及证书中的主机名和服务器主机名是否一致。
		//InsecureSkipVerify: true,

		// Set root CAs.
		RootCAs: certPool,
	}

	conn, err := tls.Dial("tcp", service, &config)
	checkError(err)
	fmt.Println("[Connected to ", conn.RemoteAddr(),"]")

	scanner := bufio.NewScanner(os.Stdin)
	var buf [512]byte
	for {
		scanner.Scan()
		_, err = conn.Write(scanner.Bytes())
		checkError(err)
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[:n]))
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
