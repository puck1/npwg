//This is a server that supports https mutual authentication.
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// You can use `ch7_security/x509Cert/genX509Cert.go` or openssl to generate certificates.
	// Private key file of this server.
	serverKeyFile := "material/server.key"
	// Cert file of this server.
	serverCertFile := "material/server.pem"
	// Root CA cert of trusted client.
	// If received client cert is not issued by this CA, connection will be refused.
	clientCAFile := "material/clientCA.pem"

	// Simplest way, use default server:
	// In this way we can't check certs delivered by client unless you modify http.DefaultServer.TLSConfig
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello, world"))		// ignore error
	//})
	//err := http.ListenAndServeTLS(":8001", serverCertFile, serverKeyFile, nil)
	//checkError(err)

	// Another way, use custom server and tls.Config:
	// In this way we can have more options such as checking client's cert.
	cert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	checkError(err)

	clientCABytes, err := ioutil.ReadFile(clientCAFile)
	checkError(err)

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(clientCABytes) {
		log.Fatal("cannot get root CA cert of client")
	}

	config := tls.Config{
		Certificates:                []tls.Certificate{cert},
		ClientAuth:                  tls.RequireAndVerifyClientCert,
		ClientCAs:                   certPool,
	}

	s := http.Server{
		Addr:              ":8001",
		// Handler specify used mux or handler, if it's nil, use http.DefaultServeMux
		Handler:           nil,
		TLSConfig:         &config,
		//ReadTimeout:       0,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))		// ignore error
	})

	// certFile is specified in Server.TLSConfig.Certificates so this can be empty.
	checkError(s.ListenAndServeTLS("", ""))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
