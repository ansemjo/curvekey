package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/nacl/box"
)

// zero and ones reader for debugging
var zero = bytes.NewBuffer(make([]byte, 128))

// exit with error when err is not nil
func fatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// print bytes in base64 or hexadecimal
func print64(data []byte) {
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)
}
func printx(data []byte) {
	fmt.Printf("%#x\n", data)
}

func main() {

	senderPub, senderSec, err := box.GenerateKey(rand.Reader)
	fatal(err)

	fmt.Println("sender")
	print64(senderSec[:])
	print64(senderPub[:])

	receiverPub, receiverSec, err := box.GenerateKey(zero)
	fatal(err)

	fmt.Println("receiver")
	print64(receiverSec[:])
	print64(receiverPub[:])

	fmt.Println("\nshared keys")

	var shared [32]byte
	box.Precompute(&shared, receiverPub, senderSec)
	print64(shared[:])

	box.Precompute(&shared, senderPub, receiverSec)
	print64(shared[:])

}
