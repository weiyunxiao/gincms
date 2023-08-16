package auth

import (
	"errors"
	"gincms/app"
	"gincms/app/common/comdata"
	"gincms/app/common/comservice"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

// AllowCurrentPath 检查是否有权限访问当前路由
func AllowCurrentPath(c *gin.Context, uid int64) (allow bool, err error) {
	user, err := comservice.FindUserByUid(c, uid)
	if err != nil {
		return false, err
	}
	if user.SuperAdmin == 1 {
		return true, nil
	}

	method := strings.ToLower(c.Request.Method)
	reqPath := strings.ToLower(c.Request.URL.Path)
	//post:/admin_api/sys/auth_login
	authStr := method + ":" + reqPath + "|"
	allThisAuthPathStr := "all" + ":" + reqPath + "|"
	/************判断是否必要权限***************/
	var paramValue string
	if paramValue, err = comservice.GetParam("needAuthButNeedAllow"); err != nil {
		return
	}
	if strings.Contains(paramValue, authStr) || strings.Contains(paramValue, allThisAuthPathStr) {
		return true, err
	}
	/************判断是否必要权限 end***************/

	//1.获取用户的角色id
	roleIdArr, err := GetUserRoleIDArr(c, uid)
	if err != nil {
		return
	}
	if len(roleIdArr) == 0 {
		return false, errors.New("暂时没有分配角色给这个用户")
	}
	//2.获取对应的菜单id
	menuIdArr, err := GetUserMenuIDArr(c, roleIdArr)
	if err != nil {
		return
	}
	if len(menuIdArr) == 0 {
		return false, errors.New("暂时没有分配菜单给这个用户")
	}
	var have int64
	err = app.DB().Model(&model.SysMenu{}).Where("deleted=0 and id in ? and (authority like ? or authority like ?)", menuIdArr, "%"+authStr+"%", "%"+allThisAuthPathStr+"%").
		Limit(1).
		Count(&have).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}

	return have > 0, nil
}

// GetUserMenu 获取用户的菜单
func GetUserMenu(c *gin.Context, uid int64, menuFields string) (list []model.SysMenu, err error) {
	list = make([]model.SysMenu, 0)
	user, err := comservice.FindUserByUid(c, uid)
	if err != nil {
		return nil, err
	}
	var menuIdArr []int
	if user.SuperAdmin != 1 {
		var roleIdArr []int
		//1.获取用户的角色id
		roleIdArr, err = GetUserRoleIDArr(c, uid)
		if err != nil {
			return
		}
		if len(roleIdArr) == 0 {
			return list, errors.New("暂时没有分配角色给这个用户")
		}
		//2.获取对应的菜单id
		menuIdArr, err = GetUserMenuIDArr(c, roleIdArr)
		if err != nil {
			return
		}
		if len(menuIdArr) == 0 {
			return list, errors.New("暂时没有分配菜单给这个用户")
		}
	}

	//3.取菜单数据
	var menuList []model.SysMenu
	if len(menuFields) == 0 {
		menuFields = "*"
	}
	query := app.DB().Select(menuFields).Where("deleted=0 and type=0")
	if len(menuIdArr) > 0 {
		query.Where("id in ?", menuIdArr)
	}
	err = query.Order("sort").Find(&menuList).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	list = comdata.TreeMenu(menuList, 0)
	return
}

// GetUserRoleIDArr 获取用户角色id数组
func GetUserRoleIDArr(c *gin.Context, uid int64) (roleIdArr []int, err error) {
	tb1 := (&model.SysRole{}).TableName()
	tb2 := (&model.SysUserM2mRole{}).TableName()
	err = app.DB().Table(tb1+" as tb1").Joins("inner join "+tb2+" as tb2 on tb1.id=tb2.role_id").
		Where("tb2.user_id=? and tb1.deleted=0", uid).
		Pluck("tb1.id", &roleIdArr).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}

// GetUserMenuIDArr 根据角色id数组获取用户的菜单id数组
func GetUserMenuIDArr(c *gin.Context, roleIDArr []int) (menuIdArr []int, err error) {
	err = app.DB().Model(&model.SysRoleM2mMenu{}).Distinct("menu_id").Where("role_id in ?", roleIDArr).Scan(&menuIdArr).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}
