package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 检查启动时参数指定的确实是hugo的站点文件夹
func CheckHugoDir(dirPath string) (valid bool) {
	isDir := IsDir(dirPath)
	if isDir {
		// 判断系统是否安装了hugo以及这个文件夹是否是hugo的文件夹
		utils.ExecCommand("hugo")
	} else {

	}
	// 子目录遍历
	// err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		log.Println("遍历子目录出错", err.Error())
	// 		return err
	// 	}
	// 	fmt.Println(path)

	// 	return nil
	// })
	// return
}

// 判断所给路径是否为文件夹  
func IsDir(path string) bool {  
	s, err := os.Stat(path)  
	if err != nil {  
			return false  
	}  
	return s.IsDir()  
}