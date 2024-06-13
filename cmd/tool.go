package cmd

import (
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"os"
)

func persistCache() {
	encoded, _ := json.Marshal(cache)
	_ = os.WriteFile(cachePath, encoded, 0666)
}

func persistEnv(offline bool, decompressed bool) {
	env.Offline = offline
	env.Decompressed = decompressed
	encoded, _ := json.Marshal(env)
	_ = os.WriteFile(envPath, encoded, 0666)
}

func homeDir() string {
	home, _ := homedir.Expand("~")
	return home
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func makeCmdDir(dir string) error {
	if _, err := os.Stat(dir); err != nil && !os.IsExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}
