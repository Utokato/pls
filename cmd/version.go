package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of pls",
	Run: func(cmd *cobra.Command, args []string) {
		doVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}

func doVersion() {
	fmt.Printf("pls version: %s, linux-command from unpkg.com version: %s\n", plsVersion, cache.GetLatestVersion())
}
