package router

import (
	"gowriter/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})

	//router.GET("/file",handler.GetFileList())
	router.GET("/fs/*path", handler.GetFileList)

	return router
}
