package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.EnableCommandSorting = false
}

var curvekey = &cobra.Command{
	Use: "curvekey",
}

func Execute() {
	err := curvekey.Execute()
	if err != nil {
		os.Exit(1)
	}
}
