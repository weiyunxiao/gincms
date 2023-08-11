package middleware

import (
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
			c.Set("uid", cast.ToInt64(jwtStruct.UserID))
			//记录排除GET的操作日志
			
		}
		c.Next()
	}
}
