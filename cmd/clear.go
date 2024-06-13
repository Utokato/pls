package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewClearCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear all info",
		Run: func(cmd *cobra.Command, args []string) {
			doClear()
		},
	}
	return cmd
}

func doClear() {
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println("[sorry] clear data fail")
	}
}
