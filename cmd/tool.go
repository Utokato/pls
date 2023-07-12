package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

func fetchAllContents(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body := resp.Body
	defer body.Close()
	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(b)
	front := "<script>window.__DATA__ = "
	back := "</script>"
	regex := regexp.MustCompile(fmt.Sprintf("%s(.*?)%s", front, back))
	target := regex.FindString(contents)
	if len(target) == 0 {
		log.Fatalln("Found not")
	}
	target = strings.Replace(target, front, "", -1)
	target = strings.Replace(target, back, "", -1)
	return target
}

func fetchAllAndCreateCache() {
	contents := fetchAllContents(fmt.Sprintf(pkgTemplate, version))
	var pkg Package
	err := json.Unmarshal([]byte(contents), &pkg)
	if err != nil {
		panic(err)
	}
	cache.LatestVersion = pkg.GetLatestVersion()
	cache.Cmds = pkg.GetCommandMaps()
}

func persistCache() {
	encoded, _ := json.Marshal(cache)
	_ = os.WriteFile(cachePath, encoded, 0666)
}

// fetchFileAndFillCache 依次发起 http 请求，将每个 command 对应的 .md 文件缓存到本地
func fetchFileAndFillCache() {
	var all, failed int64
	ch := make(chan *Cmd, 4)
	cmds := cache.GetCmds()
	wg := sync.WaitGroup{}

	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range ch {
				if err := item.FillSelf(cmdTemplate, cache.GetLatestVersion()); err != nil {
					atomic.AddInt64(&failed, 1)
				}
				atomic.AddInt64(&all, 1)
				fmt.Printf("[busy working] upgrade command:<%d/%d> => %s\n", atomic.LoadInt64(&all), len(cmds), item.Name)
			}
		}()
	}

	for _, item := range cmds {
		ch <- item
	}
	close(ch)
	wg.Wait()
	fmt.Printf("[clap] all commands are upgraded. All: %d, Failed: %d\n", all, failed)
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
