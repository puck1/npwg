package main

import (
	"crypto/x509"
	"encoding/pem"

	//"encoding/pem"
	"fmt"
	//"io"
	"io/ioutil"
	"os"
)

func main() {
	// Read from pem.
	certPEMFile, err := os.Open("localhost.pem")
	checkError(err)
	pembytes, err := ioutil.ReadAll(certPEMFile) // If ReadAll reads all, the error returned will be nil but not EOF.
	checkError(err)
	certPEMFile.Close()

	block, _ := pem.Decode(pembytes)

	cert, err := x509.ParseCertificate(block.Bytes)
	checkError(err)

	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())

	// ----------------------------------------------------------
	// Read from a common file.
	certCerFile, err := os.Open("localhost.cer")
	checkError(err)
	derBytes := make([]byte, 1000) // bigger than the file
	count, err := certCerFile.Read(derBytes)
	checkError(err)
	certCerFile.Close()

	// trim the bytes to actual length in call
	cert, err = x509.ParseCertificate(derBytes[0:count])
	checkError(err)

	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error:", err)
		os.Exit(1)
	}
}
