package file

import (
	"errors"
	"gowriter/config"
	"gowriter/utils"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DescFile struct {
	FileName    string
	FileSize    int64
	FileModTime time.Time
	FileIsDir   bool
	FilePath    string
}

// 判断所给路径是否为文件夹

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断路径是否有效

func IsValidUrl(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	} else {
		return true
	}
}

// 获取指定目录下的所有文件, 用于展示 请求url /fs/ 相对于启动时设置的站点根目录 考虑返回信息结构体包含类型、大小、创建时间

func ListAllFiles(rPath string) (descFiles []DescFile) {
	path := filepath.Join(config.GetConfig().SitePath, rPath)
	if fileDescs, err := ioutil.ReadDir(path); err != nil {
		log.Println("读取目录列表出错")
		return
	} else {
		for i := range fileDescs {
			descFiles = append(descFiles, DescFile{
				FileName:    fileDescs[i].Name(),
				FileIsDir:   fileDescs[i].IsDir(),
				FileModTime: fileDescs[i].ModTime(),
				FileSize:    fileDescs[i].Size(),
				FilePath:    strings.ReplaceAll(filepath.Join(rPath, fileDescs[i].Name()),"\\","/"),
			})
		}
	}
	return
}

// 获取文件

// 文件列表分页

// 如果要求的页数大于实际拥有的页数，返回错误

func FileListPagination(fileList []DescFile, page int, size int) (needList []DescFile, pageNum int, err error) {
	if size >= len(fileList) {
		needList = fileList
		pageNum = 1
	} else {
		pageNum = int(math.Ceil(float64(len(fileList)) / float64(size)))
		sliceStart := (page - 1) * size
		sliceEnd := sliceStart + size
		if sliceEnd >= len(fileList) {
			err = errors.New("页码越界")
		} else {
			needList = fileList[sliceStart:sliceEnd]
		}
	}
	return
}

// 检查启动时参数指定的确实是hugo的站点文件夹

func CheckHugoDir(dirPath string) (err error) {
	// 判断系统是否安装了hugo以及这个文件夹是否是hugo的文件夹
	_, stderr, err := utils.ExecCommandInPath("hugo", dirPath, "list", "all")
	if err != nil {
		log.Println("命令执行失败", stderr)
		err = errors.New("exec command failed:" + stderr)
	} else if stderr != "" {
		log.Println("hugo命令不存在")
		err = errors.New("hugo not exist" + stderr)
	} else {
		// 检查设置的目录是否在hugo站点文件夹下
		//log.Println(stdout)
		log.Println("指定目录有效")
	}

	return
}

func CheckSiteDir() (err error) {
	sitePath := config.GetConfig().SitePath
	isDir := IsDir(sitePath)
	if isDir {
		switch config.GetConfig().SiteType {
		case "hugo":
			err = CheckHugoDir(sitePath)
		}
	} else {
		// 启动参数错误，非目录
		log.Println("参数非目录或目录不存在")
		err = errors.New(sitePath + " not a directory")
	}
	return
}

// 读取文件内容

func GetContent(baseDir string, path string) (content string, err error) {
	log.Println("尝试读取", filepath.Join(baseDir,path))
	file, err := os.Open(filepath.Join(baseDir,path))
	if err != nil {
		return
	}
	defer file.Close()
	byteContent, err := ioutil.ReadAll(file)
	content = string(byteContent)
	return
}
