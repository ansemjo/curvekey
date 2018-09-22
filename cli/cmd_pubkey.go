package cli

import (
	"fmt"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

var pkout *FileFlag
var secret *Key32Flag

func init() {
	this := pubkeyCommand
	curvekey.AddCommand(this)
	this.Flags().SortFlags = false

	secret = AddKey32Flag(this, Key32FlagOptions{"key", "k", "secret key (default: stdin)", true})

	pkout = AddFileFlag(this, FileFlagOptions{
		Flag: "pub", Short: "p", Usage: "output public key to file (default: stdout)",
		Open: func(name string) (*os.File, error) {
			return os.Create(name)
		},
	})
}

var pubkeyCommand = &cobra.Command{
	Use:     "pubkey",
	Aliases: []string{"pk", "pub"},
	Short:   "Generate public key from secret key.",
	Long: `Read a secret key and perform a basepoint multiplication on Curve25519 to
calculate the public key.`,
	Example: `  curvekey pk < key.sec
  curvekey pubkey -k key.sec -p key.pub`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAll(cmd, secret.Check, pkout.Check)
	},
	Run: func(cmd *cobra.Command, args []string) {

		if pkout.File == nil {
			pkout.File = os.Stdout
		}
		defer pkout.File.Close()

		pub := keymgr.Pubkey(secret.Key)
		keymgr.Shred(secret.Key)
		fmt.Fprintln(pkout.File, encode(pub))

	},
}
