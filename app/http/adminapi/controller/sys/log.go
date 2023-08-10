package sys

import (
	"gincms/app/common/typescom"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/http/adminapi/types"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
)

var LogCtl = new(logCtl)

type logCtl struct {
}

// LoginLogoutPage 登录登出分页数据列表
func (l *logCtl) LoginLogoutPage(c *gin.Context) {
	var req types.LogLoginLogoutPageReq
	if err := c.ShouldBind(&req); err != nil {
		jsonresp.JsonFailParame(c, err)
		return
	}
	total, list, err := sys.LogService.LoginLogoutPage(c, &req)
	if err != nil {
		jsonresp.JsonFailWithMessage(err.Error(), c)
		return
	}

	jsonresp.JsonOkWithData(typescom.PageDataResp{
		Total: total,
		List:  list,
	}, c)
}
