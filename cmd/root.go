package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	plsVersion = "0.0.1"
	dir        = ".commands"
)

const (
	version     = "1.14.0"
	rootUrl     = `https://unpkg.com/linux-command@%s`
	pkgTemplate = rootUrl + "/command/"
	cmdTemplate = rootUrl + "%s"
)

var (
	dirPath   = filepath.Join(homeDir(), dir)
	cachePath = filepath.Join(homeDir(), dir, ".cache")
	root      = &cobra.Command{Use: "pls", Short: "Impressive Linux commands cheat sheet cli"}
	cache     = new(Cache)
)

// Execute all api entry.
func Execute() {
	if err := root.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// 创建根目录
	if !fileExist(dirPath) {
		err := makeCmdDir(dirPath)
		if err != nil {
			panic(err)
		}
	}
	if fileExist(cachePath) {
		parseCmdCache()
	} else {
		fmt.Println("[busy working] Building cache...")
		fetchAllAndCreateCache()
		fetchFileAndFillCache()
		persistCache()
	}
	// 构建 cobra 命令
	root.AddCommand(
		NewShowCommand(),
		NewUpgradeCommand(),
		NewVersionCommand(),
		NewSearchCommand(),
	)
}

func parseCmdCache() {
	file, _ := os.ReadFile(cachePath)
	_ = json.Unmarshal(file, cache)
}
