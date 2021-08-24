package router

import (
	"gowriter/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/assets","./assets")
	router.LoadHTMLGlob("template/**/*")
	// TODO index之后做成主页 未登录重定向至登录页
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home/index.tmpl",gin.H{
			"title": "Hello Gin",
		})
	})
	router.GET("/editor", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home/editor.tmpl",gin.H{})
	})
	router.GET("/filelist", func(context *gin.Context) {
		context.HTML(http.StatusOK,"home/posts.tmpl",gin.H{})
	})
	//router.GET("/file",handler.GetFileList())
	router.GET("/fs/*path", handler.GetFileHandler)
	// 打开编辑器的前端页面，通过api获取文章内容(文章路径通过参数给出 )
	// 前端相关的路由
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/filelist", handler.GetListHandler)
	}
	return router
}
