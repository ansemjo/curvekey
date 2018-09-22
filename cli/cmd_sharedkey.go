package cli

import (
	"fmt"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

var peerkey, mykey *Key32Flag
var sharedkey, ephemeralpub *FileFlag

func init() {
	this := sharedkeyCommand
	curvekey.AddCommand(this)
	this.Flags().SortFlags = false

	peerkey = AddKey32Flag(this, Key32FlagOptions{"peer", "p", "peer's public key (default: stdin)", true})
	mykey = AddKey32Flag(this, Key32FlagOptions{"key", "k", "your secret key (default: random)", false})

	sharedkey = AddFileFlag(this, FileFlagOptions{"shared", "s", "write shared key (default: stdout)",
		func(name string) (*os.File, error) {
			return os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		},
	})

	ephemeralpub = AddFileFlag(this, FileFlagOptions{"ephemeral", "e", "write ephemeral public key (default: stdout)",
		func(name string) (*os.File, error) {
			return os.Create(name)
		},
	})
}

var sharedkeyCommand = &cobra.Command{
	Use:     "shared",
	Aliases: []string{"dh"},
	Short:   "Agree on a shared key with a peer.",
	Long: `Read the peer's public key (and your own secret key) and agree on a shared
secret by essentially performing elliptic-curve Diffie-Hellmann on Curve25519.

If your secret is not given, an ephemeral key is created and you'll need to
transmit the ephemeral public key to your peer. This is then a trapdoor, as you
have no way to recalculate the shared secret without the peer's private key.`,
	Example: `  curvekey dh < peer.pub
  curvekey shared --peer peer.pub -e ephemeral.pub
  curvekey shared --peer ephemeral.pub --key peer.sec`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return checkAll(cmd, peerkey.Check, mykey.Check, sharedkey.Check, ephemeralpub.Check)
	},
	Run: func(cmd *cobra.Command, args []string) {

		shared, public := keymgr.SharedKey(peerkey.Key, mykey.Key)

		if sharedkey.File == nil {
			sharedkey.File = os.Stdout
			fmt.Fprint(os.Stderr, "shared secret:\n  ")
		}
		defer sharedkey.File.Close()
		fmt.Fprintln(sharedkey.File, encode(shared[:]))

		if public != nil {
			if ephemeralpub.File == nil {
				ephemeralpub.File = os.Stdout
				fmt.Fprint(os.Stderr, "ephemeral public key:\n  ")
			}
			fmt.Fprintln(ephemeralpub.File, encode(public[:]))
		}

	},
}
