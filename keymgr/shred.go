package keymgr

import "crypto/rand"

// Shred overwrites the contents of an *array with
// random data and then sets the pointer to nil
// in an attempt to clear secret data from memory
func Shred(b *[32]byte) {
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	b = nil
}
