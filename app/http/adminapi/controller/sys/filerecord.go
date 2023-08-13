package sys

import (
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
)

// 附件上传记录控制器
var FileRecordCtl = new(fileRecordCtl)

type fileRecordCtl struct {
}

// AttachmentPage 附件上传记录分页列表
func (f *fileRecordCtl) AttachmentPage(c *gin.Context) {
	var req types.AttachmentPageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.FileRecordService.AttachmentPage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}

// DelAttachment 删除
func (f *fileRecordCtl) DelAttachment(c *gin.Context) {
	var req typescom.IDArrReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}

	err := sys.FileRecordService.DelAttachment(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}
	jsonresp.JsonOk(c)
}
