package cli

import (
	"fmt"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

var peerkey **[32]byte
var peerkeyCheck checkfunc

func init() {
	this := sharedkeyCommand
	curvekey.AddCommand(this)

	peerkey, peerkeyCheck = keyFlag(this, "peer", "p", "peer's public key", true)

}

var sharedkeyCommand = &cobra.Command{
	Use:     "dh",
	Short:   "read peer's public key and output ephemeral shared key",
	Example: "  cat peerkey | curvekey dh",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return peerkeyCheck(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("%v: %v\n", peerkey, *peerkey)

		shared, public := keymgr.SharedKey(*peerkey, nil)

		fmt.Fprint(os.Stderr, "shared secret : ")
		fmt.Println(encode(shared[:]))
		if public != nil {
			fmt.Fprint(os.Stderr, "ephemeral pub : ")
			fmt.Println(encode(public[:]))
		}

	},
}
