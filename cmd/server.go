package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"net/http"
	"os"
	"path"
	"pls/offline"
	"pls/resp"
	"pls/util"
	"strings"
)

const port = 10321

func NewServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "serve",
		Short:  "Serve a web server",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			doServe()
		},
	}
	return cmd
}

func doServe() {
	if !isPortAvailable(port) {
		fmt.Println("[sorry] port is not available")
		return
	}

	// 释放静态资源
	err := util.UnarchivedTarGz(offline.Dist, ".")
	if err != nil {
		fmt.Println("[sorry] failed to extract dist.tar.gz")
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleStaticServer)
	mux.HandleFunc("/v1/healthz", handleHealthz)
	mux.HandleFunc("/v1/command/search", handleCommandSearch)
	mux.HandleFunc("/v1/command/show", handleCommandShow)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: enableCORS(mux),
	}

	// Start server
	fmt.Printf("HTTP server is running on :%d\n", port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("[sorry] failed to start server")
	}
}

// CORS 中间件
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置 CORS 响应头
		w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 处理 OPTIONS 请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleStaticServer(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./dist"))
	tmp := "./dist" + r.URL.Path
	_, err := os.Stat(tmp)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, "./dist/index.html")
		return
	}
	fs.ServeHTTP(w, r)
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	resp.Write(w, resp.Success("Ok"))
}

type SearchItem struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func handleCommandSearch(w http.ResponseWriter, r *http.Request) {
	var res []SearchItem
	// 获取参数
	key := r.URL.Query().Get("keyword")
	if key == "" {
		resp.Write(w, resp.Success(res))
		return
	}
	// 搜索
	key = strings.ToLower(key)
	for k, v := range cache.GetCmds() {
		k = strings.ToLower(k)
		if strings.Contains(k, key) {
			res = append(res, SearchItem{
				Name: v.Name,
				Desc: v.Desc,
			})
			continue
		}
		desc := strings.ToLower(v.Desc)
		if strings.Contains(desc, key) {
			res = append(res, SearchItem{
				Name: v.Name,
				Desc: v.Desc,
			})
			continue
		}
	}
	resp.Write(w, resp.Success(res))
}

func handleCommandShow(w http.ResponseWriter, r *http.Request) {
	cmds := cache.GetCmds()
	// 获取参数
	cmdName := r.URL.Query().Get("keyword")
	command, exist := cmds[cmdName]
	if !exist {
		msg := fmt.Sprintf("[sorry] could not found command <%s>\n", cmdName)
		resp.Write(w, resp.Failure(msg))
		return
	}
	// 将 .md 文件内容返回给界面
	fp := path.Join(dirPath, fmt.Sprintf("%s.md", command.Name))
	if !fileExist(fp) {
		msg := fmt.Sprintf("[sorry] could not found command <%s>\n", cmdName)
		resp.Write(w, resp.Failure(msg))
		return
	}
	source, err := os.ReadFile(fp)
	if err != nil {
		msg := fmt.Sprintf("[sorry] failed to open file <%s>\n", fp)
		resp.Write(w, resp.Failure(msg))
		return
	}
	resp.Write(w, resp.Success(string(source)))
}

func isPortAvailable(port int) bool {
	// 尝试监听指定端口
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		// 端口不可用
		return false
	}
	// 释放端口
	_ = ln.Close()
	// 端口可用
	return true
}
