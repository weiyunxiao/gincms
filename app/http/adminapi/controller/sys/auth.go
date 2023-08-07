package sys

import (
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
)

var AuthCtl = new(authCtl)

type authCtl struct {
}

// Logout 退出操作
func (a *authCtl) Logout(c *gin.Context) {
	//todo 将用户token加入黑名单
	jsonresp.JsonOk(c)
}
