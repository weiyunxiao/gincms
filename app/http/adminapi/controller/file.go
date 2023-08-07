package controller

import (
	"gincms/app"
	"gincms/app/http/adminapi/service"
	"gincms/pkg/jsonresp"

	"github.com/gin-gonic/gin"
)

var FileCtl = new(fileCtl)

type fileCtl struct {
}

// DirList 得到目录列表
func (f *fileCtl) DirList(c *gin.Context) {
	var fileManage service.FileManageService = &service.LocalFileManageService{}
	list, err := fileManage.DirList(app.Config.App.UploadDir)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(list, c)
}

// DirAndFileList 得到目录列表
func (f *fileCtl) DirAndFileList(c *gin.Context) {
	var fileManage service.FileManageService = &service.LocalFileManageService{}
	list, err := fileManage.DirAndFileList(app.Config.App.UploadDir)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(list, c)
}
