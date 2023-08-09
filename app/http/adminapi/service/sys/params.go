package sys

import (
	"errors"
	"gincms/app"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ParamsService = new(paramsService)

type paramsService struct {
}

// AddParams 添加操作
func (p *paramsService) AddParams(c *gin.Context, req *types.ParamsAddSaveReq) (err error) {
	addData := &model.SysParams{
		ParamName:  req.ParamName,
		ParamType:  req.ParamType,
		ParamKey:   req.ParamKey,
		ParamValue: req.ParamValue,
		Remark:     req.Remark,
		Deleted:    0,
		Creator:    c.GetInt64("uid"),
		CreateTime: carbon.DateTime{
			Carbon: carbon.Now(),
		},
		Updater: c.GetInt64("uid"),
		UpdateTime: carbon.DateTime{
			Carbon: carbon.Now(),
		},
	}
	err = app.DB().Create(&addData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// GetParam 获取单条信息
func (p *paramsService) GetParam(c *gin.Context, req *typescom.IDReq) (one model.SysParams, err error) {
	one = model.SysParams{}
	err = app.DB().Where("id=?", req.ID).Take(&one).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if one.ID == 0 {
		err = errors.New("无法找到记录")
		return
	}
	return
}

// UpdateParams 更新
func (p *paramsService) UpdateParams(c *gin.Context, req *types.ParamsAddSaveReq) (err error) {
	var oldData model.SysParams
	err = app.DB().Where("id=?", req.Id).Take(&oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if oldData.ID == 0 {
		return errors.New("未找到要修改的记录")
	}
	oldData.ParamName = req.ParamName
	oldData.ParamType = req.ParamType
	oldData.ParamKey = req.ParamKey
	oldData.ParamValue = req.ParamValue
	oldData.Remark = req.Remark
	oldData.Updater = c.GetInt64("uid")
	oldData.UpdateTime = carbon.DateTime{Carbon: carbon.Now()}
	err = app.DB().Where("id=?", req.Id).Select("*").Updates(&oldData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}

// DelParams 删除
func (p *paramsService) DelParams(c *gin.Context, req *typescom.IDArrReq) (err error) {
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysParams{}).Where("id in ?", req.IDArr).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// ParamsPage 列表分页
func (d *paramsService) ParamsPage(c *gin.Context, req *types.ParamsPageReq) (total int64, list []model.SysParams, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc, "id")
	list = make([]model.SysParams, 0)

	err = app.DB().Model(&model.SysParams{}).Where("deleted=0").Count(&total).Select("*").Order(sortStr).Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
