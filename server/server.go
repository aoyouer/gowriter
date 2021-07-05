package server

import (
	"gowriter/router"
)


func Start() {
	router := router.InitRouter()
	router.Run("127.0.0.1:8080")
}
