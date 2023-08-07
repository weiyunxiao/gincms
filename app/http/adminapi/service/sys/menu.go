package sys

import (
	"gincms/app"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var MenuService = new(menuService)

type menuService struct {
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
