package cmd

import (
	"encoding/json"
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
	home, _ := os.UserHomeDir()
	return home
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func makeCmdDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}
