package main

import (
	"flag"
	"fmt"
	"gowriter/file"
	"gowriter/model"
	"gowriter/server"
)

var (
	config model.Config
)

func init() {
	flag.StringVar(&config.SitePath, "d", "/var/www/html", "path of hugo website")
	flag.StringVar(&config.SiteType, "m", "hugo", "site type (hugo or hexo)")
}

func main() {
	flag.Parse()
	if err := file.CheckHugoDir(config.SitePath); err != nil {
		fmt.Println("启动参数检查失败", err.Error())
	} else {
		server.Start()
	}
}
