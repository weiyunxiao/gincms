package router

import (
	"gincms/app"
	"gincms/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.Static("upload", "upload") //上传文件都可以访问

	r.GET("ping", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Use(middleware.UUID())
	if app.Config.App.IsDemo {
		r.Use(middleware.CheckDemoEnv())
	}
	/************后台基础管理系统***************/
	route := r.Group("admin_api")
	routeNeedJwt := r.Group("admin_api", middleware.JWTCheck()) //需要jwt验证
	UseOperateMiddle(routeNeedJwt)
	/************后台基础管理系统 end***************/
	//不是演示环境，并且开启了记录非get操作日志记录
	AdminApiRouter(route, routeNeedJwt)
}

// UseOperateMiddle 是否使用日志记录非get操作
func UseOperateMiddle(routeNeedJwt *gin.RouterGroup) *gin.RouterGroup {
	if !app.Config.App.IsDemo && app.Config.App.IsOpenRecordOperate {
		routeNeedJwt.Use(middleware.OperationRecord())
	}
	return routeNeedJwt
}
