package middleware

import (
	"gincms/app"
	"gincms/pkg/jsonresp"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"strings"
)

// CheckDemoEnv 演示展示不可以更改数据
func CheckDemoEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		if app.Config.App.IsDemo && c.Request.Method != "GET" {
			method := strings.ToLower(c.Request.Method)
			reqPath := strings.ToLower(c.Request.URL.Path)
			isAllow := lo.ContainsBy(app.Config.App.IsDemoWhiteList, func(item string) bool {
				return strings.ToLower(item) == method+":"+reqPath
			})
			if !isAllow {
				jsonresp.JsonFailWithMessage("演示环境无法进行此操作", c)
				c.Abort()
				return
			}
		}
	}
}
