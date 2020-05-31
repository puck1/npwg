// This is a client that supports https mutual authentication.
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], "https://host:port/page")
		os.Exit(1)
	}
	service := os.Args[1]

	// You can use `ch7_security/x509Cert/genX509Cert.go` or openssl to generate certificates.
	// Private key file of this client.
	clientKeyFile := "material/client.key"
	// Cert file of this client.
	clientCertFile := "material/client.pem"
	// Root CA cert of trusted server.
	// If received server cert is not issued by this CA, connection will be refused.
	serverCAFile := "material/serverCA.pem"

	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	checkError(err)

	serverCABytes, err := ioutil.ReadFile(serverCAFile)
	checkError(err)

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCABytes) {
		log.Fatal("Cannot get rootCA cert of server")
	}

	config := tls.Config{
		// InsecureSkipVerify用来控制客户端是否证书和服务器主机名。
		// 如果设置为true,则不会校验证书以及证书中的主机名和服务器主机名是否一致。
		//InsecureSkipVerify: true,

		// Server root CAs.
		RootCAs: certPool,
		// Certs of this client to send to server.
		Certificates: []tls.Certificate{clientCert},
	}

	transport := http.Transport{
		TLSClientConfig: &config,
	}
	// Transport is a field of http.Client to handle transport layer such as setting proxy or tls config.
	client := http.Client{
		Transport:     &transport,
	}

	req, err := http.NewRequest("GET", service, nil)
	checkError(err)
	res, err := client.Do(req)
	checkError(err)
	b, err := ioutil.ReadAll(res.Body)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(string(b))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
