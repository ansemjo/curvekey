package cli

import (
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

type key32 *[32]byte
type checkfunc func(cmd *cobra.Command) error

func keyFlag(cmd *cobra.Command, flag, shortflag, usage string, stdin bool) (keyptr **[32]byte, check checkfunc) {

	str := cmd.Flags().StringP(flag, shortflag, "", usage)
	var key *[32]byte

	return &key, func(cmd *cobra.Command) (err error) {

		changed := cmd.Flag(flag).Changed

		if !changed {

			if stdin {
				key, err = decodeKeyFile(os.Stdin)
			}

		} else {

			if is32ByteBase64Encoded(*str) {

				key, err = decodeKey([]byte(*str))

			} else {

				var file *os.File
				file, err = os.Open(*str)
				if err != nil {
					return
				}
				defer file.Close()

				key, err = decodeKeyFile(file)

			}

		}

		return

	}

}

// check if a string is a base64 encoded 32 byte value
var is32ByteBase64EncodedRegexp = regexp.MustCompile("^[A-Za-z0-9+/]{43}=$")

func is32ByteBase64Encoded(str string) bool {
	return is32ByteBase64EncodedRegexp.MatchString(str)
}
