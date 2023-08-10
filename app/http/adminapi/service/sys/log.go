package sys

import (
	"gincms/app"
	"gincms/app/http/adminapi/types"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

var LogService = new(logService)

type logService struct {
}

// LoginLogoutPage 登录登出列表分页
func (d *logService) LoginLogoutPage(c *gin.Context, req *types.LogLoginLogoutPageReq) (total int64, list []model.SysLogLogin, err error) {
	sortStr := pkg.SortStr(req.Order, req.Asc, "id")
	list = make([]model.SysLogLogin, 0)
	query := app.DB().Model(&model.SysLogLogin{})
	if len(req.Username) > 0 {
		query.Where("username like ?", "%"+req.Username+"%")
	}
	if len(req.Address) > 0 {
		query.Where("address like ?", "%"+req.Address+"%")
	}
	if len(req.Status) > 0 {
		query.Where("status=?", cast.ToInt8(req.Status))
	}
	err = query.Count(&total).Select("*").Order(sortStr).Find(&list).Error
	if err != nil {
		app.Logger.Error("sql错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
	}
	return
}
