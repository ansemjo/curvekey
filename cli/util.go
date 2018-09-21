package cli

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func fatal(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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
