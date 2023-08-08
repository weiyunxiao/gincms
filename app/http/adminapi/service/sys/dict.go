package sys

import (
	"gincms/app"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var DictService = new(dictService)

type dictService struct {
}

// TypePage 分页数据
func (u *dictService) TypePage(c *gin.Context, req *types.DictPageReq) (total int64, list []model.SysDictType, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc, "id")
	list = make([]model.SysDictType, 0)

	query := app.DB().Model(&model.SysDictType{}).Where("deleted=0")
	if len(req.DictName) > 0 {
		query.Where("dict_name like ?", "%"+req.DictName+"%")
	}
	if len(req.DictType) > 0 {
		query.Where("dict_type like ?", "%"+req.DictType+"%")
	}

	err = query.Count(&total).Select("*").Order(sortStr).
		Scopes(pkg.PaginateScope(req.Page, req.Limit)).
		Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// DictTypeAll 获取所有字典数据
func (m *dictService) DictTypeAll(c *gin.Context) (dictTypes []model.SysDictType, err error) {
	dictTypes = make([]model.SysDictType, 0)
	err = app.DB().Select("id,dict_type").Where("deleted=0").
		Preload("DataList", func(db *gorm.DB) *gorm.DB {
			return db.Select("dict_type_id,dict_label,dict_value,label_class").Where("deleted=0").Order("sort")
		}).Find(&dictTypes).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
