package handler

import (
	"github.com/gin-gonic/gin"
	"gowriter/config"
	"gowriter/file"
	"log"
	"path/filepath"
)

// 获取指定目录下的文件列表，如果指定的是一个文件则返回文件内容

func GetFileHandler(context *gin.Context) {
	basePath := config.GetConfig().SitePath
	subPath := context.Param("path")
	destPath := filepath.Join(basePath, subPath)
	log.Println("尝试访问", destPath)
	if file.IsValidUrl(destPath) {
		if file.IsDir(destPath) {
			// 当路径是目录的时候展示目录内容
			fileDescs := file.ListAllFiles(destPath)
			context.JSON(200, fileDescs)
		} else {
			// 当路径是文件的时候 返回文件
			context.Header("Content-Description", "File Transfer")
			context.Header("Content-Transfer-Encoding", "binary")
			context.Header("Content-Disposition", "attachment; filename="+filepath.Base(destPath))
			context.Header("Content-Type", "application/octet-stream")
			context.File(destPath)
		}

	} else {
		// 无效路径 404
		log.Printf("访问路径%s不存在", destPath)
	}
}
