package main

import (
	"flag"
	"github.com/Kirov7/FayGateway/dao"
	"github.com/Kirov7/FayGateway/http_proxy_router"
	"github.com/Kirov7/FayGateway/router"
	"github.com/Kirov7/FayGateway/tcp_proxy_router"
	"github.com/Kirov7/go-config/lib"
	"os"
	"os/signal"
	"syscall"
)

// endpoint dashboard后台管理 server代理服务器
// conf ./conf/prod/ 对应配置文件夹

var (
	endpoint = flag.String("endpoint", "dashboard", "input endpoint dashboard or server")
	config   = flag.String("conf", "./conf/dev/", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()

	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *endpoint == "dashboard" {
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		router.HttpServerStop()
	} else {
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		dao.ServiceManagerHandler.LoadOnce()
		dao.AppManagerHandler.LoadOnce()
		go func() {
			http_proxy_router.HttpServerRun()
		}()
		go func() {
			http_proxy_router.HttpsServerRun()
		}()
		go func() {
			tcp_proxy_router.TcpServerRun()
		}()
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		tcp_proxy_router.TcpServerStop()
		http_proxy_router.HttpsServerStop()
		http_proxy_router.HttpServerStop()
	}

}
