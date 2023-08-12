package filepkg

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"mime/multipart"
	"path/filepath"
	"time"
)

type FileManageServiceInterface interface {
	DirList(parentDir string) ([]map[string]any, error)                //获取目录
	DirAndFileList(parentDir string) ([]map[string]interface{}, error) //获取目录及文件
	UploadFile(c *gin.Context, file *multipart.FileHeader) (filePath string, fileName string, size int64, sizeTip string, err error)
}

// CreateFileManage 工厂模式创建对应实例
func CreateFileManage(fileType string) FileManageServiceInterface {
	switch fileType {
	case "local":
		return &LocalFileManage{}
	}
	return nil
}

// createRandomFileName 生成随机的文件名
func createRandomFileName(file *multipart.FileHeader) string {
	return cast.ToString(time.Now().Unix()) + lo.RandomString(10, lo.NumbersCharset) + filepath.Ext(file.Filename)
}
