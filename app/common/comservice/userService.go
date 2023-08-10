package comservice

import (
	"errors"
	"gincms/app"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// FindUserByUuid 根据uid获取用户信息
func FindUserByUuid(c *gin.Context, uuid int64) (user *model.SysUser, err error) {
	var u model.SysUser
	err = app.DB().Where("id=?", uuid).Take(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	if u.ID == 0 {
		err = errors.New("用户不存在")
		return nil, err
	}
	return &u, nil
}

// FindUserNameByUuid 根据uuid获取用户名
func FindUserNameByUuid(c *gin.Context) (userName string) {
	uuid := c.GetInt64("uid")
	err := app.DB().Model(&model.SysUser{}).Select("username").Where("id=?", uuid).Limit(1).Scan(&userName).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		return
	}
	return
}
