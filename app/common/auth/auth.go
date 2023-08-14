package auth

import (
	"gincms/app"
	"gincms/app/common/comdata"
	"gincms/app/common/comservice"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUserMenu 获取用户的菜单
func GetUserMenu(c *gin.Context, uid int64, menuFields string) (list []model.SysMenu, err error) {
	user, err := comservice.FindUserByUid(c, uid)
	if err != nil {
		return nil, err
	}
	var menuIdArr []int
	if user.SuperAdmin != 1 {
		var roleIdArr []int
		//1.获取用户的角色id
		err = app.DB().Model(&model.SysUserM2mRole{}).Select("role_id").Where("user_id=?", uid).Scan(&roleIdArr).Error
		if err != nil {
			app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
			return
		}
		//2.获取对应的菜单id
		err = app.DB().Model(model.SysRoleM2mMenu{}).Distinct("menu_id").Where("role_id in ?", roleIdArr).Scan(&menuIdArr).Error
		if err != nil {
			app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
			return
		}
		if len(menuIdArr) == 0 {
			menuIdArr = []int{0}
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
