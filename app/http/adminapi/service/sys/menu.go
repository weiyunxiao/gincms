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

var MenuService = new(menuService)

type menuService struct {
}

// AddMenu 添加
func (m *menuService) AddMenu(c *gin.Context, req *types.MenuAddSaveReq) (err error) {
	var addData model.SysMenu

	addData.Type = req.Type
	addData.Name = req.Name
	addData.Pid = req.Pid
	addData.URL = req.Url
	addData.Authority = req.Authority
	addData.Sort = req.Sort
	addData.Icon = req.Icon
	addData.OpenStyle = req.OpenStyle
	addData.Creator = c.GetInt64("uid")
	addData.CreateTime = carbon.DateTime{Carbon: carbon.Now()}
	addData.Updater = c.GetInt64("uid")
	addData.UpdateTime = carbon.DateTime{Carbon: carbon.Now()}
	err = app.DB().Create(&addData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// GetMenu 获取单条信息
func (m *menuService) GetMenu(c *gin.Context, req *typescom.IDReq) (one model.SysMenu, err error) {
	one = model.SysMenu{}
	err = app.DB().Where("id=?", req.ID).Preload("ParentMenu", "deleted=0").Take(&one).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if one.ID == 0 {
		err = errors.New("无法找到记录")
		return
	}
	if one.ParentMenu != nil {
		one.ParentName = one.ParentMenu.ParentName
	}
	return
}

// UpdateMenu 更新单条
func (m *menuService) UpdateMenu(c *gin.Context, req *types.MenuAddSaveReq) (err error) {
	var oldData model.SysMenu
	err = app.DB().Where("id=?", req.Id).Take(&oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if oldData.ID == 0 {
		return errors.New("未找到要修改的记录")
	}
	oldData.Type = req.Type
	oldData.Name = req.Name
	oldData.Pid = req.Pid
	oldData.URL = req.Url
	oldData.Authority = req.Authority
	oldData.Sort = req.Sort
	oldData.Icon = req.Icon
	oldData.OpenStyle = req.OpenStyle
	oldData.Updater = c.GetInt64("uid")
	oldData.UpdateTime = carbon.DateTime{Carbon: carbon.Now()}
	err = app.DB().Where("id=?", req.Id).Select("*").Updates(&oldData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}

	return
}

// MenuList 系统所有菜单-树结构列表
func (m *menuService) MenuList(c *gin.Context, typeParam int) (menuList []model.SysMenu, err error) {
	menuList = make([]model.SysMenu, 0)
	selectFields := "id,pid,name,url,type,open_style,icon,authority,sort,create_time"
	err = app.DB().Select(selectFields).Where("deleted=0 and type=?", typeParam).Order("sort").Find(&menuList).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}

	tree := comdata.TreeMenu(menuList, 0)
	return tree, nil
}

// DelMenu 删除单条
func (r *menuService) DelMenu(c *gin.Context, req *typescom.IDReq) (err error) {
	var have int64
	err = app.DB().Model(&model.SysMenu{}).Where("pid=? and deleted=0", req.ID).Limit(1).Count(&have).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if have > 0 {
		err = errors.New("有未删除的子菜单，无法删除该菜单")
		return
	}
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysMenu{}).Where("id= ?", req.ID).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}

// Nav 查询用户菜单
func (m *menuService) Nav(c *gin.Context) (menuList []model.SysMenu, err error) {
	menuList = make([]model.SysMenu, 0)
	selectFields := "id,pid,name,url,icon,sort,create_time"
	err = app.DB().Select(selectFields).Where("deleted=0 and pid=0").Order("sort").Preload("Children", func(db *gorm.DB) *gorm.DB {
		return db.Select(selectFields).Where("deleted=0").Order("sort")
	}).Find(&menuList).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
