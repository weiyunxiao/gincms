package sys

import (
	"context"
	"errors"
	"gincms/app"
	"gincms/app/common/comdata"
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	lop "github.com/samber/lo/parallel"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

var RoleService = new(roleService)

type roleService struct {
}

// AddRole 添加
func (u *roleService) AddRole(c *gin.Context, req *types.RoleAddSaveReq) (err error) {
	menuArr := make([]model.SysMenu, 0)
	for _, v := range req.MenuIdList {
		menuArr = append(menuArr, model.SysMenu{
			ID: v,
		})
	}
	var addData model.SysRole

	addData.Name = req.Name
	addData.Remark = req.Remark
	addData.Creator = c.GetInt64("uid")
	addData.RoleMenuList = menuArr
	addData.CreateTime = carbon.DateTime{Carbon: carbon.Now()}
	addData.Updater = c.GetInt64("uid")
	addData.UpdateTime = carbon.DateTime{Carbon: carbon.Now()}
	err = app.DB().Omit("RoleMenuList.*").Create(&addData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// UpdateRole 更新
func (u *roleService) UpdateRole(c *gin.Context, req *types.RoleAddSaveReq) (err error) {
	var oldData model.SysRole
	err = app.DB().Where("id=?", req.Id).Take(&oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if oldData.ID == 0 {
		return errors.New("未找到要修改的记录")
	}
	var count int64
	err = app.DB().Model(&model.SysRole{}).Where("name=? and id <> ?", req.Name, req.Id).Limit(1).Count(&count).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if count > 0 {
		return errors.New("系统中已存在这个名称了")
	}

	oldData.Name = req.Name
	oldData.Remark = req.Remark
	oldData.Updater = c.GetInt64("uid")
	oldData.UpdateTime = carbon.DateTime{Carbon: carbon.Now()}

	menuArr := make([]model.SysMenu, 0, len(req.MenuIdList))
	if len(req.MenuIdList) > 0 {
		err = app.DB().Where("id in ?", req.MenuIdList).Find(&menuArr).Error
		if err != nil {
			app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
			return
		}
	}

	err = app.DB().Transaction(func(tx *gorm.DB) error {
		err = app.DB().Where("id=?", req.Id).Select("*").Updates(&oldData).Error
		if err != nil {
			return err
		}
		/************更新拥有菜单信息***************/
		err = app.DB().Model(&oldData).Association("RoleMenuList").Clear()
		if err != nil {
			return err
		}
		if len(menuArr) > 0 {
			err = app.DB().Model(&oldData).Omit("RoleMenuList.*").Association("RoleMenuList").Append(menuArr)
			if err != nil {
				return err
			}
		}
		/************更新拥有菜单信息 end***************/
		return nil
	})
	if err != nil {
		app.Logger.Error("更新角色错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}

	return
}

// GetRole 获取一条角色信息
func (r *roleService) GetRole(c *gin.Context, req *typescom.IDReq) (one model.SysRole, err error) {
	one = model.SysRole{}
	err = app.DB().Where("id=?", req.ID).Preload("RoleMenuList", "deleted=0").Take(&one).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if one.ID == 0 {
		err = errors.New("无法找到记录")
		return
	}
	one.MenuIdList = make([]int, 0, len(one.RoleMenuList))
	if len(one.RoleMenuList) > 0 {
		var locker sync.Mutex
		lop.ForEach(one.RoleMenuList, func(item model.SysMenu, _ int) {
			locker.Lock()
			defer locker.Unlock()
			one.MenuIdList = append(one.MenuIdList, int(item.ID))
		})
		one.RoleMenuList = nil
	}
	return
}

// RolePage 分页数据
func (r *roleService) RolePage(c *gin.Context, req *types.RolePageReq) (total int64, list []model.SysRole, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc, "id")
	list = make([]model.SysRole, 0)

	query := app.DB().Model(&model.SysRole{}).Where("deleted=0")
	if len(req.Name) > 0 {
		query.Where("name like ?", "%"+req.Name+"%")
	}

	err = query.Count(&total).Select("*").Order(sortStr).
		Scopes(pkg.PaginateScope(req.Page, req.Limit)).Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// DelRole 根据id数组删除
func (r *roleService) DelRole(c *gin.Context, req *typescom.IDArrReq) (err error) {
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysRole{}).Where("id in ?", req.IDArr).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// RoleList 获取可用角色列表数据
func (r *roleService) RoleList(c *gin.Context) (list []model.SysRole, err error) {
	sortStr := pkg.SortStr("", false, "id")
	list = make([]model.SysRole, 0)

	query := app.DB().Model(&model.SysRole{}).Where("deleted=0")
	err = query.Select("id,name").Order(sortStr).
		Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}

// RoleUserPage 某个角色拥有的用户分页列表
func (r *roleService) RoleUserPage(c *gin.Context, req *types.RoleHaveUserPageReq) (total int64, list []model.SysUser, err error) {
	list = make([]model.SysUser, 0)
	var role model.SysRole
	err = app.DB().Where("id=?", req.RoleId).Take(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if role.ID == 0 {
		err = errors.New("未找到这个角色信息")
		return
	}

	/************取分页数据***************/
	query := app.DB().Model(&role)
	if len(req.Username) > 0 {
		query.Where("username like ?", "%"+req.Username+"%")
	}
	if len(req.Mobile) > 0 {
		query.Where("mobile like ?", "%"+req.Mobile+"%")
	}
	if len(req.Gender) > 0 {
		query.Where("gender=?", cast.ToInt(req.Gender))
	}
	query2 := query.WithContext(context.Background())
	total = query.Association("RoleUserList").Count()

	err = query2.Scopes(pkg.PaginateScope(req.Page, req.Limit)).Association("RoleUserList").Find(&list)
	if err != nil {
		app.Logger.Error("错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	/************取分页数据 end***************/
	return
}

// DelRoleUser 移除某个角色下的用户
func (r *roleService) DelRoleUser(c *gin.Context, roleId int, req *typescom.IDArrReq) (err error) {
	err = app.DB().Where("role_id=? and user_id in ?", roleId, req.IDArr).Delete(&model.SysUserM2mRole{}).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}

// AddRoleUser 某个角色关联多个用户的操作
func (r *roleService) AddRoleUser(c *gin.Context, roleId int, req *typescom.IDArrReq) (err error) {
	var role model.SysRole
	err = app.DB().Where("id=?", roleId).Take(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if role.ID == 0 {
		err = errors.New("请指定正确的角色")
		return
	}

	var addDAta = make([]model.SysUserM2mRole, 0)
	for _, uid := range req.IDArr {
		data := model.SysUserM2mRole{
			UserID: uid,
			RoleID: int64(roleId),
		}
		addDAta = append(addDAta, data)
	}
	err = app.DB().Clauses(clause.OnConflict{DoNothing: true}).Create(&addDAta).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}

// RoleMenu 添加编辑角色时，显示的所有系统菜单
func (r *roleService) RoleMenu(c *gin.Context) (tree []model.SysMenu, err error) {
	allData := make([]model.SysMenu, 0)
	sortStr := pkg.SortStr("", true)
	err = app.DB().Where("deleted=0").Order(sortStr).Find(&allData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	tree = comdata.TreeMenu(allData, 0)
	return
}
