package keymgr

import (
	"golang.org/x/crypto/curve25519"
)

// Pubkey takes a secret key and perform a scalar base
// multiplication to obtain the public key
func Pubkey(seckey *[32]byte) *[32]byte {
	pubkey := new([32]byte)
	curve25519.ScalarBaseMult(pubkey, seckey)
	return pubkey
}
