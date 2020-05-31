/* Package blowfish implements Bruce Schneier's Blowfish encryption algorithm.

Blowfish is a legacy cipher and its short block size makes it vulnerable to
birthday bound attacks (see https://sweet32.info). It should only be used where
compatibility with legacy systems, not security, is the goal.

Deprecated: any new system should use AES (from crypto/aes, if necessary in an AEAD mode like crypto/cipher.NewGCM)
or XChaCha20-Poly1305 (from golang.org/x/crypto/chacha20poly1305).

It is just a demo aimed to show how to use symmetric cipher, etc.
 */
package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/blowfish"
	"os"
)

func main() {
	key := []byte("my key")
	cipher, err := blowfish.NewCipher(key)
	checkError(err)

	src := []byte("some words more than 8 bytes")
	// Need to pad it with extra blanks to make sure it's aligned to the block size, 8.
	ensureAligned(&src)

	// Can only encrypt 8 bytes once.
	var encryBuf [8]byte
	encrypted := bytes.NewBuffer(nil)

	for i := 0; i < len(src); i += 8 {
		cipher.Encrypt(encryBuf[0:], src[i:i+8])
		encrypted.Write(encryBuf[0:])
	}

	transmited := encrypted.Bytes()

	// Can only decrypt 8 bytes once.
	var decryBuf [8]byte
	decrypted := bytes.NewBuffer(nil)
	for i := 0; i < len(src); i += 8 {
		// Other endpoints can use the key to instantiate a CIPHER for their use.
		cipher.Decrypt(decryBuf[0:], transmited[i:i+8])
		decrypted.Write(decryBuf[0:])
	}
	fmt.Println(decrypted)
}

func ensureAligned(text *[]byte) {
	n := len(*text) % 8
	if n == 0 {
		return
	}
	for i := 0; i < 8-n; i++ {
		*text = append(*text, ' ')
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
