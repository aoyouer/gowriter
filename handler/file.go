package handler

import (
	"github.com/gin-gonic/gin"
	"gowriter/config"
	"log"
	"path/filepath"
)

// 获取指定目录下的文件列表，如果指定的是一个文件则返回文件内容

func GetFileList(context *gin.Context) {
	basePath := config.GetConfig().SitePath
	subPath := context.Param("path")
	destPath := filepath.Join(basePath,subPath)
	log.Println("尝试访问",destPath)
}
