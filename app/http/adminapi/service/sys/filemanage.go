package sys

import (
	"gincms/app"
	"gincms/app/common/typescom"
	"gincms/app/model"
	"gincms/pkg"
	"gincms/pkg/filepkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

var FileManageService = new(fileManageService)

type fileManageService struct {
}

// UploadFile 上传文件
func (f *fileManageService) UploadFile(c *gin.Context, req *typescom.FileUploadReq) (filePath string, fileName string, err error) {
	var fileManage = filepkg.CreateFileManage()
	filePath, fileName, size, sizeTip, err := fileManage.UploadFile(c, req.File)
	if err != nil {
		return
	}
	var addData = model.SysAttachment{
		Name:       fileName,
		URL:        filePath,
		Size:       size,
		SizeTip:    sizeTip,
		Platform:   app.Config.App.OssType,
		TenantID:   0,
		Version:    0,
		Deleted:    0,
		Creator:    c.GetInt64("uid"),
		CreateTime: carbon.DateTime{Carbon: carbon.Now()},
		Updater:    c.GetInt64("uid"),
		UpdateTime: carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Create(&addData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
