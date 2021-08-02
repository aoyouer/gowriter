package server

import (
	"gowriter/router"
)


func Start() {
	r := router.InitRouter()
	r.Run("0.0.0.0:8080")
}
