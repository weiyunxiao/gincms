package filepkg

import (
	"fmt"
	"gincms/app"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/goutil/fmtutil"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"mime/multipart"
	"os"
)

type LocalFileManage struct{}

// DirList 获取目录
func (f *LocalFileManage) DirList(parentDir string) (dirs []map[string]any, err error) {
	dirs = make([]map[string]any, 0)
	files, err := os.ReadDir(parentDir)
	if err != nil {
		return
	}
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			dir := make(map[string]any, 0)
			dir["id"] = files[i].Name()
			dir["label"] = files[i].Name()
			dirs = append(dirs, dir)
		}
	}
	return
}

// DirAndFileList 获取目录及文件
func (f *LocalFileManage) DirAndFileList(parentDir string) (files []map[string]interface{}, err error) {
	files = make([]map[string]interface{}, 0)
	fList, err := os.ReadDir(parentDir)
	if err != nil {
		return
	}
	webUrl := app.Config.Http.WebSiteUrl
	for i := 0; i < len(fList); i++ {
		info := make(map[string]interface{})
		info["id"] = i   //前端配合用
		info["url"] = "" //前端

		info["is_dir"] = false
		info["fileName"] = fList[i].Name()
		info["size"] = ""

		if fList[i].IsDir() {
			info["is_dir"] = true
		} else {
			f, errF := fList[i].Info()
			if errF != nil {
				err = fmt.Errorf("读取文件%s,发生错误:%v", fList[i].Name(), errF)
				return
			}
			info["url"] = webUrl + "/" + parentDir + "/" + fList[i].Name() //前端
			info["size"] = fmtutil.DataSize(uint64(f.Size()))
		}
		files = append(files, info)
	}
	return
}

// UploadFile 本地上传文件
func (f *LocalFileManage) UploadFile(c *gin.Context, file *multipart.FileHeader) (filePath string, fileName string, size int64, sizeTip string, err error) {
	// 确保目录存在，不存在创建
	dirName := app.Config.App.UploadDir + "/" + carbon.Now().ToDateString() + "/"
	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		app.Logger.Error("创建本地上传目录时出错:", zap.Error(err))
		return
	}
	fileName = createRandomFileName(file)
	filePath = dirName + fileName
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		app.Logger.Error("保存文件失败:", zap.Error(err))
		return
	}
	size = file.Size
	sizeTip = fmtutil.SizeToString(cast.ToUint64(size))
	return
}
