package cmd

import (
	"github.com/spf13/cobra"
)

func NewUpgradeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade all commands from remote.",
		Run: func(cmd *cobra.Command, args []string) {
			doUpgrade()
		},
	}
	return cmd
}

func doUpgrade() {
	// 获取 command 文件目录
	fetchAllAndCreateCache()
	// 获取每一个 command 文件
	fetchFileAndFillCache()
	// 持久化 cache
	persistCache()
}
