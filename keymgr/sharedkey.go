// Copyright (c) 2018 Anton Semjonov
// Licensed under the MIT License

package keymgr

import (
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/salsa20/salsa"
)

// SharedKey takes a peer's public key and our own secret key to
// obtain a shared secret through elliptic-curve Diffie-Hellmann.
// If mySecret is nil, a random secret key will be generated and
// the ephemeral public key will be output in myPublic. The peer
// needs our public key to be able to obtain the shared secret.
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
