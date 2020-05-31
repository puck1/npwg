package main

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	var key rsa.PrivateKey
	loadKey("private.key", &key)

	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	var pubKey rsa.PublicKey
	loadKey("public.key", &pubKey)


	fmt.Println("Public key modulus", pubKey.N.String())
	fmt.Println("Public key exponent", pubKey.E)
}

func loadKey(file string, key interface{}) {
	fd, err := os.Open(file)
	checkError(err)

	decoder := gob.NewDecoder(fd)
	err = decoder.Decode(key)
	checkError(err)

	err = fd.Close()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
