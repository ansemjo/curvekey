package cli

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

func readkey(file *os.File) *[32]byte {

	key, err := ioutil.ReadAll(file)
	fatal(err)

	n, err := base64.StdEncoding.Decode(key, key)
	fatal(err)

	if n != 32 {
		fatal(errors.New("key must be 32 bytes"))
	}

	return get32(key)

}
