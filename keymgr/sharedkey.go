package keymgr

import (
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/salsa20/salsa"
)

func SharedKey(peerPublic, mySecret *[32]byte) (shared, myPublic *[32]byte) {

	shared = new([32]byte)
	if mySecret == nil {
		mySecret = NewKey()
		myPublic = Pubkey(mySecret)
	}

	// kdf compatible with box.Precompute
	curve25519.ScalarMult(shared, mySecret, peerPublic)
	salsa.HSalsa20(shared, new([16]byte), shared, &salsa.Sigma)

	return

}
