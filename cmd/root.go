package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	plsVersion = "1.0.1"
	dir        = ".commands"
)

const (
	version     = "1.18.0"
	rootUrl     = `https://unpkg.com/linux-command@%s`
	pkgTemplate = rootUrl + "/command/"
	cmdTemplate = rootUrl + "%s"
)

var (
	dirPath   = filepath.Join(homeDir(), dir)
	envPath   = filepath.Join(homeDir(), dir, ".env")
	cachePath = filepath.Join(homeDir(), dir, ".cache")

	root  = &cobra.Command{Use: "pls", Short: "Impressive Linux commands cheat sheet cli"}
	env   = new(Env)
	cache = new(Cache)
)

// Execute all api entry.
func Execute() {
	if err := root.Execute(); err != nil {
		// ignore error
		// _ = root.Help()
	}
}

func init() {
	// 创建根目录
	if !fileExist(dirPath) {
		err := makeCmdDir(dirPath)
		if err != nil {
			fmt.Println("[sorry] failed to create command dir")
		}
	}

	if fileExist(envPath) {
		parseEnv()
	} else {
		fmt.Println("[tips] env info is not found, so setting it to online mode.")
		setDefaultEnv()
	}

	if fileExist(cachePath) {
		parseCache()
	} else {
		fmt.Println("[tips] cache info is not found, please use offline cmd to unzip resource or use upgrade cmd to update resource.")
	}
	// 构建 cobra 命令
	root.AddCommand(
		NewOfflineCommand(),
		NewShowCommand(),
		NewUpgradeCommand(),
		NewVersionCommand(),
		NewSearchCommand(),
		NewClearCommand(),
		NewServeCommand(),
	)
}

func setDefaultEnv() {
	persistEnv(false, false)
}

func parseCache() {
	file, _ := os.ReadFile(cachePath)
	_ = json.Unmarshal(file, cache)
}

func parseEnv() {
	file, _ := os.ReadFile(envPath)
	_ = json.Unmarshal(file, env)
}
