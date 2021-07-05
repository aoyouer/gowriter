package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 检查启动时参数指定的确实是hugo的站点文件夹
func CheckHugoDir(dir string) (err error) {
	// 子目录遍历
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("遍历子目录出错", err.Error())
			return err
		}
		fmt.Println(path)
		// if info.IsDir() {
		// 	// 忽略目录
		// }
		return nil
	})
	return
}
