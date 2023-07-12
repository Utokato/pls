package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version of pls",
		Run: func(cmd *cobra.Command, args []string) {
			doVersion()
		},
	}
	return cmd
}

func doVersion() {
	fmt.Printf("pls version: %s, linux-command from unpkg.com version: %s\n", plsVersion, cache.GetLatestVersion())
}
