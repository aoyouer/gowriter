package router

import (
	"gowriter/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(config model.Config) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	router.StaticFS("/fs", http.Dir(config.SitePath))

	return router
}
