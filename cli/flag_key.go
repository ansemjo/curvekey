package cli

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

type Key32 *[32]byte
type CmdCheckFunc func(cmd *cobra.Command) error
type Key32Flag struct {
	Key   *Key32
	Check CmdCheckFunc
}

// AddKey32Flag adds a flag to a command, which can either be a valid base64
// string or a filename for a 32 byte key. Optionally reads from stdin.
func AddKey32Flag(cmd *cobra.Command, flag, shortflag, usage string, stdin bool) Key32Flag {

	// add flag to command
	str := cmd.Flags().StringP(flag, shortflag, "", usage)
	var key Key32

	// return struct and build check function inline
	return Key32Flag{&key, func(cmd *cobra.Command) (err error) {

		// if flag was given
		if cmd.Flag(flag).Changed {

			// and it is a valid base64 encoded key
			if Is32ByteBase64Encoded(*str) {
				key, err = decodeKey([]byte(*str))

			} else {
				// assume any other string to be a filename
				file, err := os.Open(*str)
				if err != nil {
					return err
				}
				defer file.Close()
				key, err = decodeKeyFile(file)
			}

		} else if stdin {
			// if flag was not given but "read from stdin" is true
			key, err = decodeKeyFile(os.Stdin)
		}

		// if neither, just return nil. the pointer to Key will remain nil!
		return
	}}
}

// Is32ByteBase64Encoded checks if the given string is a base64-encoded 32 byte value.
func Is32ByteBase64Encoded(str string) bool {
	return regexp.MustCompile("^[A-Za-z0-9+/]{43}=$").MatchString(str)
}

// decodeKey decodes a base64 string and expects a 32 byte value inside
func decodeKey(b []byte) (key *[32]byte, err error) {
	n, err := base64.StdEncoding.Decode(b, b)
	if err != nil {
		return
	}
	if n != 32 {
		err = errors.New("key must be 32 bytes")
		return
	}
	key = get32(b)
	return
}

// decodeKeyFile reads a file and decodes its contents with decodeKey
func decodeKeyFile(file *os.File) (key *[32]byte, err error) {
	keyslice, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return decodeKey(keyslice)
}
