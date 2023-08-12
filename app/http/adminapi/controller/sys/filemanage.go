package sys

import (
	"gincms/app"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/pkg/filepkg"
	"gincms/pkg/jsonresp"
	"github.com/gookit/goutil/fsutil"
	"strings"

	"github.com/gin-gonic/gin"
)

var FileManageCtl = new(fileManageCtl)

type fileManageCtl struct {
}

// UploadFile 上传文件
func (f *fileManageCtl) UploadFile(c *gin.Context) {
	var req typescom.FileUploadReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	path, name, err := sys.FileManageService.UploadFile(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithDetailed(map[string]any{
		"filePath": path,
		"fileName": name,
	}, "操作成功", c)
}

// DownFile 下载文件
func (f *fileManageCtl) DownFile(c *gin.Context) {
	filePath := c.Query("filePath")
	if len(filePath) == 0 {
		jsonresp.JsonFailWithMessage("参数错误", c)
		return
	}
	//打开文件
	if !fsutil.IsFile(filePath) {
		jsonresp.JsonFailWithMessage("文件不存在", c)
		return
	}

	strArr := strings.Split(filePath, "/")
	fileName := strArr[len(strArr)-1]
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
	return
}

// DirList 目录列表
func (f *fileManageCtl) DirList(c *gin.Context) {
	fileType := app.Config.App.OssType
	baseDir := app.Config.App.UploadDir
	var fileManage = filepkg.CreateFileManage(fileType)
	list, err := fileManage.DirList(baseDir)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(list, c)
}

func (f *fileManageCtl) DirAndFileList(c *gin.Context) {
	fileType := app.Config.App.OssType
	baseDir := app.Config.App.UploadDir
	var fileManage = filepkg.CreateFileManage(fileType)
	list, err := fileManage.DirAndFileList(baseDir)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(list, c)
}
