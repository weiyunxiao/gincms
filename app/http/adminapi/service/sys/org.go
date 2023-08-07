package sys

import (
	"errors"
	"gincms/app"
	"gincms/app/common/comdata"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var OrgService = new(orgService)

type orgService struct {
}

// AddOrg 添加机构操作
func (p *orgService) AddOrg(c *gin.Context, req *types.OrgAddSaveReq) (err error) {
	addData := &model.SysOrg{
		Pid:     req.Pid,
		Name:    req.Name,
		Sort:    req.Sort,
		Creator: c.GetInt64("uid"),
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

// UpdateOrg 更新机构操作
func (p *orgService) UpdateOrg(c *gin.Context, req *types.OrgAddSaveReq) (err error) {
	var updateData = map[string]any{
		"name":    req.Name,
		"pid":     req.Pid,
		"sort":    req.Sort,
		"updater": c.GetInt64("uid"),
		"update_time": carbon.DateTime{
			Carbon: carbon.Now(),
		},
	}
	err = app.DB().Model(&model.SysOrg{}).Where("id=?", req.Id).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// GetOrg 查询一条机构
func (p *orgService) GetOrg(c *gin.Context, req *typescom.IDReq) (one *model.SysOrg, err error) {
	one = &model.SysOrg{}
	err = app.DB().Select("id,name,pid,sort").Where("id=?", req.ID).Preload("ParenOrg").Take(&one).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if one.ParenOrg != nil {
		one.ParentName = one.ParenOrg.Name
	}
	return
}

// DelOrg 根据id删除机构
func (p *orgService) DelOrg(c *gin.Context, req *typescom.IDReq) (err error) {
	var count int64
	err = app.DB().Model(&model.SysOrg{}).Where("pid = ? and deleted=0", req.ID).Limit(1).Count(&count).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if count > 0 {
		return errors.New("有子机构没有被删除，请先删除子机构")
	}
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysOrg{}).Where("id = ?", req.ID).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// OrgList 获取可用机构列表数据
func (o *orgService) OrgList(c *gin.Context, req typescom.PageOrderCommonReq) (listTree []model.SysOrg, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc)
	listTree = make([]model.SysOrg, 0)
	req.Page = 1
	req.Limit = 2000
	query := app.DB().Model(&model.SysOrg{}).Where("deleted=0")
	err = query.Select("id,pid,name,sort").Order(sortStr).
		Preload("ParenOrg").
		Scopes(pkg.PaginateScope(req.Page, req.Limit)).
		Find(&listTree).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return comdata.TreeOrg(listTree, 0), nil
}
