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

func decodeKeyFile(file *os.File) (key *[32]byte, err error) {
	keyslice, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return decodeKey(keyslice)
}
