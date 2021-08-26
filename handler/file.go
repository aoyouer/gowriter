package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowriter/blog"
	"gowriter/config"
	"gowriter/file"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
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

// 获取文件列表的处理,分为获取页面文件、文章文件 分别在不同的目录中进行搜索

func GetListHandler(context *gin.Context) {
	var listType, category string
	var page, pageSize int
	listType = context.DefaultQuery("type", "")
	category = context.DefaultQuery("category", "")
	page, _ = strconv.Atoi(context.DefaultQuery("page", "0"))
	pageSize, _ = strconv.Atoi(context.DefaultQuery("size", "10"))
	// 在指定文件夹进行查找并进行分页
	var fileList []file.DescFile
	switch config.GetConfig().SiteType {
	case "hugo":
		fmt.Printf("TYPE:%s Category:%s Page:%d Size:%d\n", listType, category, page, pageSize)
		fileList = blog.GetHugoFile(listType, category)
	default:
		fileList = []file.DescFile{}
	}
	if list, pageNum, err := file.FileListPagination(fileList, page, pageSize); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"currentpage": page, "totalpage": pageNum, "pagesize": pageSize, "files": list})
	}
}
