package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

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

	pubA, secA, err := box.GenerateKey(rand.Reader)
	fatal(err)

	fmt.Println("\npeer A")
	printx(secA[:])
	printx(pubA[:])

	pubB, secB, err := box.GenerateKey(rand.Reader)
	fatal(err)

	fmt.Println("\npeer B")
	printx(secB[:])
	printx(pubB[:])

	fmt.Println("\nshared keys")

	var shared [32]byte
	box.Precompute(&shared, pubB, secA)
	printx(shared[:])

	box.Precompute(&shared, pubA, secB)
	printx(shared[:])

}
