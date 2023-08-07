package service

import (
	"fmt"
	"github.com/gookit/goutil/fmtutil"
	"os"
)

type FileManageService interface {
	DirList(parentDir string) ([]string, error)                        //获取目录
	DirAndFileList(parentDir string) ([]map[string]interface{}, error) //获取目录及文件
}

type LocalFileManageService struct{}

// DirList 获取目录
func (f *LocalFileManageService) DirList(parentDir string) (dirs []string, err error) {
	dirs = make([]string, 0)
	files, err := os.ReadDir(parentDir)
	if err != nil {
		return
	}
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			dirs = append(dirs, files[i].Name())
		}
	}
	return
}

// DirAndFileList 获取目录及文件
func (f *LocalFileManageService) DirAndFileList(parentDir string) (files []map[string]interface{}, err error) {
	files = make([]map[string]interface{}, 0)
	fList, err := os.ReadDir(parentDir)
	if err != nil {
		return
	}
	for i := 0; i < len(fList); i++ {
		info := make(map[string]interface{})
		info["is_dir"] = false
		info["name"] = fList[i].Name()
		info["size"] = ""

		if fList[i].IsDir() {
			info["is_dir"] = true
		} else {
			f, errF := fList[i].Info()
			if errF != nil {
				err = fmt.Errorf("读取文件%s,发生错误:%v", fList[i].Name(), errF)
				return
			}
			info["size"] = fmtutil.DataSize(uint64(f.Size()))
		}
		files = append(files, info)
	}
	return
}
