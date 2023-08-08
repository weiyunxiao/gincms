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

var DictService = new(dictService)

type dictService struct {
}

// AddDictType 添加操作
func (d *dictService) AddDictType(c *gin.Context, req *types.DictTypeAddSaveReq) (err error) {
	addData := &model.SysDictType{
		DictType:   req.DictType,
		DictName:   req.DictName,
		Sort:       req.Sort,
		DictSource: req.DictSource,
		DictSql:    req.DictSql,
		Remark:     req.Remark,
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

// GetDictType 获取单条信息
func (d *dictService) GetDictType(c *gin.Context, req *typescom.IDReq) (one model.SysDictType, err error) {
	one = model.SysDictType{}
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

// UpdateDictType 更新单条
func (d *dictService) UpdateDictType(c *gin.Context, req *types.DictTypeAddSaveReq) (err error) {
	var oldData model.SysDictType
	err = app.DB().Where("id=?", req.Id).Take(&oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if oldData.ID == 0 {
		return errors.New("未找到要修改的记录")
	}
	oldData.DictType = req.DictType
	oldData.DictName = req.DictName
	oldData.Sort = req.Sort
	oldData.DictSource = req.DictSource
	oldData.DictSql = req.DictSql
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

// TypePage 分页数据
func (d *dictService) TypePage(c *gin.Context, req *types.DictPageReq) (total int64, list []model.SysDictType, err error) {
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

// DelType 删除多条字典类型
func (d *dictService) DelType(c *gin.Context, req *typescom.IDArrReq) (err error) {
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysDictType{}).Where("id in ?", req.IDArr).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// DictTypeAll 获取所有字典数据
func (d *dictService) DictTypeAll(c *gin.Context) (dictTypes []model.SysDictType, err error) {
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

// DictDataPage 获取字典二级数据列表分页
func (d *dictService) DictDataPage(c *gin.Context, req *types.DictDataPageReq) (total int64, list []model.SysDictData, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc)
	list = make([]model.SysDictData, 0)

	err = app.DB().Model(&model.SysDictData{}).Where("deleted=0 and dict_type_id=?", req.DictTypeId).Count(&total).Select("*").Order(sortStr).
		Scopes(pkg.PaginateScope(req.Page, req.Limit)).
		Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// DelDictData 删除多条二级字典数
func (d *dictService) DelDictData(c *gin.Context, req *typescom.IDArrReq) (err error) {
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysDictData{}).Where("id in ?", req.IDArr).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// AddDictData 添加操作
func (d *dictService) AddDictData(c *gin.Context, req *types.DictDataAddSaveReq) (err error) {
	addData := &model.SysDictData{
		DictTypeID: int64(req.DictTypeId),
		DictLabel:  req.DictLabel,
		DictValue:  req.DictValue,
		LabelClass: req.LabelClass,
		Sort:       req.Sort,
		Remark:     req.Remark,
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

// GetDictData 获取单条信息
func (d *dictService) GetDictData(c *gin.Context, req *typescom.IDReq) (one model.SysDictData, err error) {
	one = model.SysDictData{}
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

// UpdateDictData 更新单条
func (d *dictService) UpdateDictData(c *gin.Context, req *types.DictDataAddSaveReq) (err error) {
	var oldData model.SysDictData
	err = app.DB().Where("id=?", req.Id).Take(&oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if oldData.ID == 0 {
		return errors.New("未找到要修改的记录")
	}

	oldData.DictTypeID = int64(req.DictTypeId)
	oldData.DictLabel = req.DictLabel
	oldData.DictValue = req.DictValue
	oldData.LabelClass = req.LabelClass
	oldData.Sort = req.Sort
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
