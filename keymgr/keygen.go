package keymgr

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/nacl/box"
)

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

type KeyPair struct {
	Public *[32]byte
	Secret *[32]byte
}

func NewKeyPair(reader io.Reader) *KeyPair {
	pub, sec, err := box.GenerateKey(reader)
	if err != nil {
		panic(err)
	}
	return &KeyPair{pub, sec}
}

func (k *KeyPair) Print() {
	print64(k.Secret[:])
	print64(k.Public[:])
}

func Shred(b *[32]byte) {
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	b = nil
}

type EphemeralKey struct {
	Public *[32]byte
	Shared *[32]byte
}

func NewEphemeralKey(pub *[32]byte) *EphemeralKey {
	sender := NewKeyPair(rand.Reader)
	shared := new([32]byte)
	eph := &EphemeralKey{Public: sender.Public, Shared: shared}
	box.Precompute(eph.Shared, pub, sender.Secret)
	Shred(sender.Secret)
	return eph
}

func (e *EphemeralKey) Print() {
	print64(e.Shared[:])
}

// func main() {

// 	fmt.Println("receiver")
// 	receiver := newKeyPair(zero)
// 	receiver.print()

// 	fmt.Println("\nephemeral key")
// 	ephem := newEphemeralKey(receiver.Public)
// 	ephem.print()

// 	fmt.Println("\nshared key from receiver secret")
// 	var shared [32]byte
// 	box.Precompute(&shared, ephem.Public, receiver.Secret)
// 	print64(shared[:])

// }
