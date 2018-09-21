package cli

import (
	"fmt"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

var peerkey, mykey Key32Flag

func init() {
	this := sharedkeyCommand
	curvekey.AddCommand(this)

	peerkey = AddKey32Flag(this, "peer", "p", "peer's public key (default: stdin)", true)
	mykey = AddKey32Flag(this, "key", "k", "your secret key", false)
}

var sharedkeyCommand = &cobra.Command{
	Use:   "dh",
	Short: "Agree on a shared key with a peer.",
	Long: `Read the peer's public key (and your own secret key) and agree on a shared
secret by essentially performing elliptic-curve diffie-hellmann.

If your secret is not given, an ephemeral key is created and you'll need to
transmit the displayed public key to your peer. This is then a trapdoor, as you
have no way to recalculate the shared secret without the peer's private key.`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return checkAll(cmd, peerkey.Check, mykey.Check)
	},
	Run: func(cmd *cobra.Command, args []string) {

		shared, public := keymgr.SharedKey(*peerkey.Key, *mykey.Key)

		fmt.Fprint(os.Stderr, "shared secret:\n  ")
		fmt.Println(encode(shared[:]))
		if public != nil {
			fmt.Fprint(os.Stderr, "ephemeral public key:\n  ")
			fmt.Println(encode(public[:]))
		}

	},
}
