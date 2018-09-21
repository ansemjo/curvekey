package keymgr

import (
	"golang.org/x/crypto/curve25519"
)

func Pubkey(seckey *[32]byte) *[32]byte {
	pubkey := new([32]byte)
	curve25519.ScalarBaseMult(pubkey, seckey)
	return pubkey
}
