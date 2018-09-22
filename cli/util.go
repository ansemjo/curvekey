package cli

import (
	"encoding/base64"

	"github.com/spf13/cobra"
)

// base64-encode 32 byte array to string
func encode(b *[32]byte) string {
	return base64.StdEncoding.EncodeToString(b[:])
}

// copy a 32 byte array from slice
func get32(slice []byte) *[32]byte {
	key := new([32]byte)
	copy(key[:], slice)
	return key
}

// run pre-run checks of cobra flags
func checkAll(cmd *cobra.Command, checker ...func(*cobra.Command) error) (err error) {
	for _, ch := range checker {
		err = ch(cmd)
		if err != nil {
			return
		}
	}
	return
}
