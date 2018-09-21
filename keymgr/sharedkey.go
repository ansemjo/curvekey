package keymgr

import (
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/salsa20/salsa"
)

func SharedKey(peer, secret *[32]byte) (shared *[32]byte) {
	// compatible with box.Precompute
	shared = new([32]byte)
	curve25519.ScalarMult(shared, secret, peer)
	salsa.HSalsa20(shared, new([16]byte), shared, &salsa.Sigma)
	return
}

func EphemeralSharedKey(peer *[32]byte) (shared, public *[32]byte) {
	secret := NewKey()
	public = Pubkey(secret)
	shared = SharedKey(peer, secret)
	Shred(secret)
	return
}
