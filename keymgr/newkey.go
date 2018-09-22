package keymgr

import (
	"crypto/rand"
	"io"
)

// NewKey generates a new random 32 byte key
func NewKey() *[32]byte {
	key := new([32]byte)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		key = nil
		panic(err)
	}
	return key
}
