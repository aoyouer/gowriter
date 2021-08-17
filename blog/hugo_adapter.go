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

func GetHugoFile(ftype string, category string) (descFiles []file.DescFile) {
	for _, d := range directories {
		if d == ftype {
			contentPath := filepath.Join(config.GetConfig().SitePath, "content", d)
			descFiles = file.ListAllFiles(contentPath)
			return
		}
	}
	return
}
