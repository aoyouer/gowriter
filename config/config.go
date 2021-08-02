package config

import (
	"flag"
	"gowriter/file"
)

// 一个配置的单例

type Config struct {
	SitePath string
	SiteType string
}

var config *Config

func init() {
	config = &Config{}
	flag.StringVar(&config.SitePath, "d", "/var/www/html", "path of hugo website")
	flag.StringVar(&config.SiteType, "m", "hugo", "site type (hugo or hexo)")
	flag.Parse()
}

func CheckConfig() (err error) {
	err = file.CheckHugoDir(config.SitePath)
	return
}

func GetConfig() *Config {
	return config
}

