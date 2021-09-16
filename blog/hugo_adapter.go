package blog

import (
	"gowriter/config"
	"gowriter/file"
	"path/filepath"
)

// 用于控制hugo的命令

// 获取Hugo各文件夹中的文件
var (
	directories = []string{"page", "posts", "post"}
)

// TODO 根据传入的分类对文章、分类进行过滤 (需要读取文章中的标签)
// TODO 还需要返回文件相对于站点根目录的路径

func GetHugoFile(ftype string, category string) (descFiles []file.DescFile) {
	for _, d := range directories {
		if d == ftype {
			contentPath := filepath.Join("content", d)
			descFiles = file.ListAllFiles(contentPath)
			return
		}
	}
	return
}

// TODO 获取文章列表，需要有标签字段、分类字段、链接、摘要

// TODO 获取文章内容

func GetHugoContent(path string) (content string, err error) {
	content, err = file.GetContent(config.GetConfig().SitePath,path)
	return
}
