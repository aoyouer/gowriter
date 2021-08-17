package server

import (
	"gowriter/config"
	"gowriter/router"
)

func Start() {
	r := router.InitRouter()
	r.Run(config.GetConfig().BindAddr)
}
