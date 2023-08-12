package sys

import (
	"gincms/app/common/comservice"
	"gincms/pkg/jsonresp"
	"gincms/pkg/jwt"
	"github.com/gin-gonic/gin"
)

var AuthCtl = new(authCtl)

type authCtl struct {
}

// RefreshToken 刷新token
func (a *authCtl) RefreshToken(c *gin.Context) {
	refreshToken := c.DefaultQuery("refreshToken", "")
	if len(refreshToken) == 0 {
		jsonresp.JsonFailParameWithMsg("请提供刷新token参数", c)
		return
	}
	jwtObj := jwt.NewJWT()
	jwtStruct, err := jwtObj.ParserToken(c, refreshToken)

	if err != nil || !jwtStruct.IsFreshToken {
		jsonresp.JsonResult(jsonresp.RefreshToken_code, gin.H{}, err.Error(), c)
		return
	}

	if !jwtStruct.IsFreshToken {
		jsonresp.JsonResult(jsonresp.RefreshToken_code, gin.H{}, "不是刷新所使用的token", c)
		return
	}
	user, err := comservice.FindUserByUid(c, int64(jwtStruct.UserID))
	if err != nil || user.Deleted == 1 {
		jsonresp.JsonResult(jsonresp.RefreshToken_code, gin.H{}, "无法生成新的token,或用户不可用", c)
		return
	}

	//重新生成token
	token, expireAtUnix := jwtObj.IssueToken(c, jwtStruct.UserID)
	//access_token
	jsonresp.JsonOkWithData(gin.H{
		"access_token":   token,
		"expire_at_unix": expireAtUnix,
	}, c)
}

// Logout 退出操作
func (a *authCtl) Logout(c *gin.Context) {
	//todo 将用户token加入黑名单
	userName := comservice.FindUserNameByContextUid(c)
	comservice.LogLoginAndLogout(c, 1, userName)
	jsonresp.JsonOk(c)
}
