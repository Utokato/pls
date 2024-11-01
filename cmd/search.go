package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var searchCommand = &cobra.Command{
	Use:   "search <command>",
	Short: "Search command by keywords",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("[sorry] the search command does not accept any keywords")
			return
		}
		doSearch(args[0])
	},
}

func init() {
	rootCmd.AddCommand(searchCommand)
}

func doSearch(keyword string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"command", "description"})
	table.SetRowLine(true)
	keyword = strings.ToLower(keyword)
	for k, v := range cache.GetCmds() {
		k = strings.ToLower(k)
		if strings.Contains(k, keyword) {
			table.Append([]string{v.Name, v.Desc})
			continue
		}
		desc := strings.ToLower(v.Desc)
		if strings.Contains(desc, keyword) {
			table.Append([]string{v.Name, v.Desc})
			continue
		}
	}
	table.Render()
}
