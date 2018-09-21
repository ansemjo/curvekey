package cli

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

func init() {
	curvekey.AddCommand(pubkeyCommand)
}

var pubkeyCommand = &cobra.Command{
	Use:     "pk",
	Short:   "read a secret key and output the public key",
	Example: "  cat secret | curvekey pk > pubkey",
	Run: func(cmd *cobra.Command, args []string) {

		keyin, err := ioutil.ReadAll(os.Stdin)
		fatal(err)

		n, err := base64.StdEncoding.Decode(keyin, keyin)
		fatal(err)

		if n != 32 {
			fatal(errors.New("key must be 32 bytes"))
		}

		pub := keymgr.Pubkey(get32(keyin))
		fmt.Println(encode(pub[:]))
	},
}
