package file

import (
	"errors"
	"gowriter/utils"
	"log"
	"os"
)

// 检查启动时参数指定的确实是hugo的站点文件夹

func CheckHugoDir(dirPath string) (err error) {
	isDir := IsDir(dirPath)
	if isDir {
		// 判断系统是否安装了hugo以及这个文件夹是否是hugo的文件夹
		stdout, stderr, err := utils.ExecCommandInPath("hugo", dirPath)
		if err != nil {
			log.Println("命令执行失败", stderr)
			err = errors.New("exec command failed:" + stderr)
		} else if stderr != "" {
			log.Println("hugo命令不存在")
			err = errors.New("hugo not exist" + stderr)
		} else {
			// 检查设置的目录是否在hugo站点文件夹下
			log.Println(stdout)
		}
	} else {
		// 启动参数错误，非目录
		log.Println("参数非目录或目录不存在")
		err = errors.New(dirPath + " not a directory")
	}
	return
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

// 获取指定目录下的所有文件, 用于展示 请求url /fs/ 相对于启动时设置的站点根目录

func ListAllFiles(path string) {

}
