package cli

import (
	"fmt"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

var secret Key32Flag

func init() {
	this := pubkeyCommand
	curvekey.AddCommand(this)
	secret = AddKey32Flag(this, "key", "k", "secret key (default: stdin)", true)
}

var pubkeyCommand = &cobra.Command{
	Use:     "pubkey",
	Aliases: []string{"pub"},
	Short:   "Generate public key from secret key.",
	Long: `Read a secret key and perform a basepoint multiplication on Curve25519 to
calculate the public key and output that on stdout.`,
	Example: "  curvekey pub -k my.sec > my.pub",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAll(cmd, secret.Check)
	},
	Run: func(cmd *cobra.Command, args []string) {

		pub := keymgr.Pubkey(*secret.Key)
		fmt.Println(encode(pub[:]))

	},
}
