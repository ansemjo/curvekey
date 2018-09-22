package cli

import (
	"fmt"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

var kgout *FileFlag
var pubout *FileFlag

func init() {
	this := newkeyCommand
	curvekey.AddCommand(this)
	this.Flags().SortFlags = false

	kgout = AddFileFlag(this, FileFlagOptions{
		Flag: "key", Short: "k", Usage: "output secret key to file (default: stdout)",
		Open: func(name string) (*os.File, error) {
			return os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		},
	})

	pubout = AddFileFlag(this, FileFlagOptions{
		Flag: "pub", Short: "p", Usage: "output public key to file (default: none)",
		Open: func(name string) (*os.File, error) {
			return os.Create(name)
		},
	})

}

var newkeyCommand = &cobra.Command{
	Use:     "keygen",
	Aliases: []string{"kg", "new"},
	Short:   "Generate a new secret key.",
	Long: `Generate a new secret key from system randomness.
If the '--pub' flag is given, the public key is calculated and written to that
file aswell. Otherwise only the secret key will be written.`,
	Example: `  curvekey new -k key.sec
  curvekey keygen -k key.sec -p key.pub`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAll(cmd, kgout.Check, pubout.Check)
	},
	Run: func(cmd *cobra.Command, args []string) {

		if kgout.File == nil {
			kgout.File = os.Stdout
		}
		defer kgout.File.Close()

		key := keymgr.NewKey()
		fmt.Fprintln(kgout.File, encode(key))

		if pubout.File != nil {
			defer pubout.File.Close()
			pub := keymgr.Pubkey(key)
			fmt.Fprintln(pubout.File, encode(pub))
		}

		keymgr.Shred(key)

	},
}
