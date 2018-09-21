package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/nacl/box"
)

// zero reader for debugging
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

	pubA, secA, err := ed25519.GenerateKey(zero)
	fatal(err)
	secA = secA[:32]

	fmt.Println("\npeer A")
	printx(secA)
	printx(pubA)

	pubB, secB, err := ed25519.GenerateKey(rand.Reader)
	fatal(err)
	secB = secB[:32]

	fmt.Println("\npeer B")
	printx(secB)
	printx(pubB)

	var sharedA [32]byte
	box.Precompute(sharedA, pubB, secA)

	var sharedB [32]byte
	box.Precompute(sharedB, pubA, secB)

	fmt.Println("\nshared keys")
	printx(sharedA)
	printx(sharedB)

}
