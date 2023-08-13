package sys

import (
	"gincms/app"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

var FileRecordService = new(fileRecordService)

type fileRecordService struct {
}

// AttachmentPage 附件上传记录列表分页
func (d *fileRecordService) AttachmentPage(c *gin.Context, req *types.AttachmentPageReq) (total int64, list []model.SysAttachment, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc, "id")
	list = make([]model.SysAttachment, 0)
	query := app.DB().Model(&model.SysAttachment{}).Where("deleted=0")
	if len(req.Name) > 0 {
		query.Where("name like ?", "%"+req.Name+"%")
	}
	if len(req.Platform) > 0 {
		query.Where("platform=?", req.Platform)
	}

	err = query.Count(&total).Select("*").Order(sortStr).
		Scopes(pkg.PaginateScope(req.Page, req.Limit)).
		Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// DelAttachment 删除
func (d *fileRecordService) DelAttachment(c *gin.Context, req *typescom.IDArrReq) (err error) {
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysAttachment{}).Where("id in ?", req.IDArr).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
