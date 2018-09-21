package keymgr

import "crypto/rand"

func Shred(b *[32]byte) {
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	b = nil
}
