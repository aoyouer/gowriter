package server

import (
	"gowriter/model"
	"gowriter/router"
)


func Start(config model.Config) {
	router := router.InitRouter(config)
	router.Run("0.0.0.1:8080")
}
