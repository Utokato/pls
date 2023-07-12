package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	ErrCommandNotFound = errors.New("command not found")
)

type Package struct {
	PackageName       string   `json:"packageName"`
	PackageVersion    string   `json:"packageVersion"`
	AvailableVersions []string `json:"availableVersions"`
	Filename          string   `json:"filename"`
	Target            Target   `json:"target"`
}

type Target struct {
	Path    string            `json:"path"`
	Type    string            `json:"type"`
	Details map[string]Detail `json:"details"`
}

type Detail struct {
	Path        string `json:"path"`
	Type        string `json:"type"`
	ContentType string `json:"contentType"`
	Integrity   string `json:"integrity"`
	Size        int    `json:"size"`
}

type Cmd struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Desc string `json:"desc"`
}

type Cache struct {
	LatestVersion string          `json:"latestVersion"`
	Cmds          map[string]*Cmd `json:"cmds"`
}

func (c *Cache) GetLatestVersion() string {
	return c.LatestVersion
}

func (c *Cache) GetCmds() map[string]*Cmd {
	return c.Cmds
}

func (pkg *Package) GetLatestVersion() string {
	return pkg.AvailableVersions[len(pkg.AvailableVersions)-1]
}

func (pkg *Package) GetCommandMaps() map[string]*Cmd {
	inner := make(map[string]*Cmd, 512)
	for k := range pkg.Target.Details {
		s := strings.Replace(k, "/command/", "", -1)
		s = strings.Replace(s, ".md", "", -1)
		inner[s] = &Cmd{
			Name: s,
			Path: k,
		}
	}
	return inner
}

// FillSelf 发起 Http 请求获取 .md 文件，将文件存储到本地中
// 并读取到每个 .md 文件中描述信息，缓存到 Cmd 的 desc 字段中
func (cmd *Cmd) FillSelf(urlTemplate, latestVersion string) error {
	url := fmt.Sprintf(urlTemplate, latestVersion, cmd.Path)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotFound {
		return ErrCommandNotFound
	}
	body := resp.Body
	defer body.Close()
	content := make([]byte, 0)
	reader := bufio.NewReader(resp.Body)
	idx := 0
	for {
		line, _, err := reader.ReadLine()
		idx++
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		content = append(content, line...)
		content = append(content, []byte("\n")...)
		// .md 文件的第四行为描述文件
		if idx == 4 {
			cmd.Desc = string(line)
		}
	}
	fp := path.Join(dirPath, fmt.Sprintf("%s.md", cmd.Name))
	err = os.WriteFile(fp, content, 0666)
	if err != nil {
		return err
	}
	return nil
}
