package cli

import (
	"fmt"
	"os"

	"github.com/ansemjo/curvekey/keymgr"
	"github.com/spf13/cobra"
)

func init() {
	curvekey.AddCommand(sharedkeyCommand)
	sharedkeyCommand.Flags().BytesBase64Var(&mysecret, "my", nil, "my secret key")
}

var mysecret []byte

var sharedkeyCommand = &cobra.Command{
	Use:     "dh",
	Short:   "read peer's public key and output ephemeral shared key",
	Example: "  cat peerkey | curvekey dh",
	Run: func(cmd *cobra.Command, args []string) {

		peerkey := readkey(os.Stdin)
		var secret, shared, public *[32]byte

		if cmd.Flag("my").Changed && len(mysecret) == 32 {
			secret = get32(mysecret)
		}

		shared, public = keymgr.SharedKey(peerkey, secret)

		fmt.Fprint(os.Stderr, "shared secret : ")
		fmt.Println(encode(shared[:]))
		if public != nil {
			fmt.Fprint(os.Stderr, "ephemeral pub : ")
			fmt.Println(encode(public[:]))
		}

	},
}
