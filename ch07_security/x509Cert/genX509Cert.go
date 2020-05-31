package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader

	keyFile := "../pubKeyEncryption/private.key"		// a private key
	var key rsa.PrivateKey
	loadKey(keyFile, &key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000) // one year in nanosecond
	//then := now.AddDate(1, 0, 0) 								// same

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "localhost",
			Organization: []string{"local"},
		},
		NotBefore: time.Now(),
		NotAfter: then,

		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage: x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA: true,
		DNSNames: []string{"localhost"},
	}

	derbytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)		// note params are pointers
	checkError(err)

	certFile, err := os.Create("localhost.cer")
	checkError(err)
	_, err = certFile.Write(derbytes)
	checkError(err)
	certFile.Close()

	certPEMFile, err := os.Create("localhost.pem")
	checkError(err)
	err = pem.Encode(certPEMFile, &pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   derbytes,
	})
	checkError(err)
	certPEMFile.Close()

	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	err = pem.Encode(keyPEMFile, &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(&key),
	})
	keyPEMFile.Close()
}

func loadKey(file string, key interface{}) {
	fd, err := os.Open(file)
	checkError(err)
	decoder := gob.NewDecoder(fd)
	err = decoder.Decode(key)
	fd.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
