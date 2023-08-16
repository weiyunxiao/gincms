package middleware

import (
	"gincms/app/common/auth"
	"gincms/pkg/jsonresp"
	"gincms/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

// JWTCheck 验证 jwt中间件
func JWTCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtObj := jwt.NewJWT()
		if jwtStruct, err := jwtObj.ParserToken(c); err != nil {
			jsonresp.JsonResult(http.StatusUnauthorized, gin.H{}, err.Error(), c)
			c.Abort()
			return
		} else {
			uid := cast.ToInt64(jwtStruct.UserID)
			c.Set("uid", uid)
			/************检查是否有权限访问***************/
			allow, err := auth.AllowCurrentPath(c, uid)
			if err != nil {
				jsonresp.JsonResult(http.StatusForbidden, gin.H{}, err.Error(), c)
				c.Abort()
				return
			}
			if !allow {
				jsonresp.JsonResult(http.StatusForbidden, gin.H{}, "没有权限访问此操作", c)
				c.Abort()
				return
			}
			/************检查是否有权限访问 end***************/
		}
		c.Next()
	}
}
