package sys

import (
	"gincms/app"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var DictService = new(dictService)

type dictService struct {
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
