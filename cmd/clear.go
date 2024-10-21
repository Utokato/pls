package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var clearCommand = &cobra.Command{
	Use:   "clear",
	Short: "Clear all info",
	Run: func(cmd *cobra.Command, args []string) {
		doClear()
	},
}

func init() {
	rootCmd.AddCommand(clearCommand)
}

func doClear() {
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println("[sorry] clear data fail")
	}
}
