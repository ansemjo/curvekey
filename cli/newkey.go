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
	Use:   "kg",
	Short: "generate a new secret key",
	Run: func(cmd *cobra.Command, args []string) {

		key := keymgr.NewKey()
		fmt.Println(encode(key[:]))

	},
}
