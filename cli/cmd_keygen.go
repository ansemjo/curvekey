package cli

import (
	"fmt"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

var out *FileFlag
var pubout *FileFlag

func init() {
	this := newkeyCommand
	curvekey.AddCommand(this)

	out = AddFileFlag(this, FileFlagOptions{
		Flag: "out", Short: "o", Usage: "output secret key to file (default: stdout)",
		Open: func(name string) (*os.File, error) {
			return os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		},
	})

	pubout = AddFileFlag(this, FileFlagOptions{
		Flag: "pubout", Short: "p", Usage: "output public key to file (default: none)",
		Open: func(name string) (*os.File, error) {
			return os.Create(name)
		},
	})

}

var newkeyCommand = &cobra.Command{
	Use:     "keygen",
	Aliases: []string{"kg", "new"},
	Short:   "Generate a new secret key.",
	Long:    `Generate a new secret key from system randomness.`,
	Example: "  curvekey new -o key.sec -p key.pub",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAll(cmd, out.Check, pubout.Check)
	},
	Run: func(cmd *cobra.Command, args []string) {

		if out.File == nil {
			out.File = os.Stdout
		}
		defer out.File.Close()

		key := keymgr.NewKey()
		fmt.Fprintln(out.File, encode(key[:]))

		if pubout.File != nil {
			defer pubout.File.Close()
			pub := keymgr.Pubkey(key)
			fmt.Fprintln(pubout.File, encode(pub[:]))
		}

	},
}
