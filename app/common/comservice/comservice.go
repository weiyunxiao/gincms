package comservice

import (
	"gincms/app"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

// GetParam 获取数据表中一个配置的参数数据
func GetParam(keyName string) (value string, ok bool) {
	err := app.DB().Model(&model.SysParams{}).Select("param_value").
		Where("param_key=? and deleted=0", keyName).
		Limit(1).Scan(&value).Error
	if err != nil {
		app.Logger.Error("sql错误222", zap.Error(err))
	}
	return value, len(value) > 0
}

// LogLoginAndLogout 登录退出日志记录
// actionType: 0:登录成功,1:退出成功,2:验证码错误
func LogLoginAndLogout(c *gin.Context, actionType int8, userName string) {
	addData := model.SysLogLogin{
		Username:   userName,
		Ip:         c.ClientIP(),
		Address:    "",
		Operation:  actionType,
		UserAgent:  c.Request.UserAgent(),
		CreateTime: carbon.DateTime{Carbon: carbon.Now()},
	}
	switch actionType {
	case 0: //登录成功
		addData.Status = 1
	case 1: //退出成功
		addData.Status = 1
	case 2: //验证码错误
		addData.Status = 0
	case 3: //账号密码错误
		addData.Status = 0
	default:
		return
	}
	err := app.DB().Create(&addData).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
}
