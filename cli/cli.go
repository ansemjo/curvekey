// Copyright (c) 2018 Anton Semjonov
// Licensed under the MIT License

package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.EnableCommandSorting = false
}

var curvekey = &cobra.Command{
	Use:     "curvekey",
	Long:    "Elliptic curve operations on Curve25519.",
	Version: "0.0.2",
}

// Execute executes the cobra cli
func Execute() {
	err := curvekey.Execute()
	if err != nil {
		os.Exit(1)
	}
}
