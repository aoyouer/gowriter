package main

import (
	"flag"
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
	file.CheckHugoDir(config.SitePath)
	server.Start()
}
