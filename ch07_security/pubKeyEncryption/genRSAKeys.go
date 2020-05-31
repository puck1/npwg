package main

import (
	"crypto/rand"		// rand.Reader
	"crypto/rsa"
	"crypto/x509"		// x509.MarshalPKCS1PrivateKey()
	"encoding/gob"
	"encoding/pem"		// pem.Block & pem.Encode()
	"fmt"
	"os"
)

func main() {
	// rand.Reader is a global, shared instance of a cryptographically.
	reader := rand.Reader
	bitSize := 2048

	// size: byte size of key pair
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes", key.Primes[0], key.Primes[1])
	fmt.Println("Private key exponent", key.D.String())

	pubKey := key.PublicKey
	fmt.Println("Public key modules", pubKey.N.String())
	fmt.Println("Public key exponent", pubKey.E)

	saveGobKey("private.key", key)
	saveGobKey("public.key", pubKey)
	savePEMKey("private.pem", key)
}

func saveGobKey(file string, key interface{}) {
	fd, err := os.Create(file)		// Even if file exists, os.Creat won't return an error.
	checkError(err)
	encoder := gob.NewEncoder(fd)
	err = encoder.Encode(key)
	checkError(err)
	fd.Close()
}

func savePEMKey(file string, key *rsa.PrivateKey) {
	fd, err := os.Create(file)		// Even if file exists, os.Creat won't return an error.
	checkError(err)

	var privateKey = &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(fd, privateKey)
	checkError(err)

	fd.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
