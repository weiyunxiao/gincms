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
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var UserService = new(userService)

type userService struct {
}

// AddUser 添加一个用户
func (u *userService) AddUser(c *gin.Context, req *types.UserAddSaveReq) (err error) {
	postArr := make([]model.SysPost, 0)
	for _, v := range req.PostIdList {
		postArr = append(postArr, model.SysPost{
			ID: v,
		})
	}
	roleArr := make([]model.SysRole, 0)
	for _, v := range req.RoleIdList {
		roleArr = append(roleArr, model.SysRole{
			ID: v,
		})
	}

	addData := &model.SysUser{
		Username:     req.Username,
		RealName:     req.RealName,
		Gender:       cast.ToInt8(req.Gender),
		Email:        req.Email,
		Mobile:       req.Mobile,
		OrgID:        req.OrgId,
		Status:       req.Status,
		Creator:      c.GetInt64("uid"),
		CreateTime:   carbon.DateTime{Carbon: carbon.Now()},
		Updater:      c.GetInt64("uid"),
		UpdateTime:   carbon.DateTime{Carbon: carbon.Now()},
		UserPostList: postArr,
		UserRoleList: roleArr,
	}
	if len(req.Password) > 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		if err != nil {
			app.Logger.Error("pwd加密出错", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
			return errors.New("无法添加")
		}
		addData.Password = string(password)
	}
	err = app.DB().Omit("UserPostList.*,UserRoleList.*").Create(&addData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// GetUser 查询一条用户
func (u *userService) GetUser(c *gin.Context, req *typescom.IDReq) (one *model.SysUser, err error) {
	one = &model.SysUser{}
	err = app.DB().Omit("password").Preload(clause.Associations).Where("id=?", req.ID).Take(&one).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if one.OrgData != nil {
		//one.OrgName = one.OrgData.Name
	}
	if len(one.UserRoleList) > 0 {
		var roleIds = make([]uint, 0, len(one.RoleIdList))
		lo.ForEach(one.UserRoleList, func(item model.SysRole, _ int) {
			roleIds = append(roleIds, item.ID)
		})
		one.RoleIdList = roleIds
	}

	if len(one.UserPostList) > 0 {
		var postIds = make([]uint, 0, len(one.PostIdList))
		lo.ForEach(one.UserPostList, func(item model.SysPost, _ int) {
			postIds = append(postIds, item.ID)
		})
		one.PostIdList = postIds
	}

	return
}

// UpdateUser 更新
func (u *userService) UpdateUser(c *gin.Context, req *types.UserAddSaveReq) (err error) {
	var oldData model.SysUser
	err = app.DB().Where("id=?", req.Id).Take(&oldData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if oldData.ID == 0 {
		return errors.New("未找到要修改的记录")
	}
	var count int64
	err = app.DB().Model(&model.SysUser{}).Where("username=? and id <> ?", req.Username, req.Id).Limit(1).Count(&count).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if count > 0 {
		return errors.New("系统中已有这个用户名了")
	}

	oldData.Username = req.Username
	oldData.RealName = req.RealName
	oldData.OrgID = req.OrgId
	oldData.Gender = cast.ToInt8(req.Gender)
	oldData.Email = req.Email
	oldData.Mobile = req.Mobile
	oldData.Status = cast.ToInt8(req.Status)
	oldData.Updater = c.GetInt64("uid")
	oldData.UpdateTime = carbon.DateTime{Carbon: carbon.Now()}

	if len(req.Password) > 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		if err != nil {
			app.Logger.Error("pwd加密出错", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
			return errors.New("生成密码出错")
		}
		oldData.Password = string(password)
	}

	roleArr := make([]model.SysRole, 0, len(req.RoleIdList))

	if len(req.RoleIdList) > 0 {
		err = app.DB().Where("id in ?", req.RoleIdList).Find(&roleArr).Error
		if err != nil {
			app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
			return
		}
	}
	//for _, v := range req.RoleIdList {
	//	roleArr = append(roleArr, model.SysRole{
	//		ID: v,
	//	})
	//}
	postArr := make([]model.SysPost, 0, len(req.PostIdList))
	if len(req.PostIdList) > 0 {
		err = app.DB().Where("id in ?", req.PostIdList).Find(&postArr).Error
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
		/************更新角色信息***************/
		err = app.DB().Model(&oldData).Association("UserRoleList").Clear()
		if err != nil {
			return err
		}
		if len(roleArr) > 0 {
			err = app.DB().Model(&oldData).Omit("UserRoleList.*").Association("UserRoleList").Append(roleArr)
			if err != nil {
				return err
			}
		}
		/************更新角色信息 end***************/

		/************更新岗位信息***************/
		err = app.DB().Model(&oldData).Association("UserPostList").Clear()
		if err != nil {
			return err
		}
		if len(postArr) > 0 {
			err = app.DB().Model(&oldData).Omit("UserPostList.*").Association("UserPostList").Append(postArr)
			if err != nil {
				return err
			}
		}
		/************更新岗位信息 end***************/
		return nil
	})
	if err != nil {
		return
	}

	return
}

// DelUser 根据id数组删除
func (u *userService) DelUser(c *gin.Context, req *typescom.IDArrReq) (err error) {
	updateData := map[string]any{
		"deleted":     1,
		"updater":     c.GetInt64("uid"),
		"update_time": carbon.DateTime{Carbon: carbon.Now()},
	}
	err = app.DB().Model(&model.SysUser{}).Where("id in ?", req.IDArr).Updates(updateData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// UserPage 分页数据
func (u *userService) UserPage(c *gin.Context, req *types.UserPageReq) (total int64, list []model.SysUser, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc, "id")
	list = make([]model.SysUser, 0)

	query := app.DB().Model(&model.SysUser{}).Where("deleted=0")
	if len(req.UserName) > 0 {
		query.Where("username like ?", "%"+req.UserName+"%")
	}
	if len(req.Mobile) > 0 {
		query.Where("mobile like ?", "%"+req.Mobile+"%")
	}
	if len(req.Gender) > 0 {
		query.Where("gender=?", cast.ToInt(req.Gender))
	}

	err = query.Count(&total).Select("id,username,real_name,avatar,gender,email,mobile,status,org_id,create_time").Order(sortStr).
		Scopes(pkg.PaginateScope(req.Page, req.Limit)).
		Preload("OrgData", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,name")
		}).Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}

// IsUserNameExist 判断用户名是否存在
func (u *userService) IsUserNameExist(c *gin.Context, username string) (isExist bool, err error) {
	var count int64
	err = app.DB().Model(&model.SysUser{}).Where("username=?", username).Limit(1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	if count > 0 {
		isExist = true
	}
	return
}

// Info 用户登录时，获取用户信息给前端
func (u *userService) Info(c *gin.Context, uid int64) (user map[string]any, err error) {
	user = make(map[string]any, 0)
	err = app.DB().Model(&model.SysUser{}).
		Select("avatar,email,gender,id,mobile,real_name as realName,status,super_admin as superAdmin,username").
		Where("id=?", uid).
		Take(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
