package cli

import (
	"fmt"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

func init() {
	curvekey.AddCommand(newkeyCommand)
}

var newkeyCommand = &cobra.Command{
	Use:     "keygen",
	Aliases: []string{"kg", "new"},
	Short:   "Generate a new secret key.",
	Long:    `Generate a new secret key from system randomness.`,
	Example: "  curvekey new > key.sec\n  curvekey pub -k key.sec > key.pub",
	Run: func(cmd *cobra.Command, args []string) {

		key := keymgr.NewKey()
		fmt.Println(encode(key[:]))

	},
}
