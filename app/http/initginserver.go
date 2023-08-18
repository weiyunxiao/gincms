package http

import (
	"gincms/app"
	"gincms/app/http/middleware"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
)

// CreateGinServer gin初始化
func CreateGinServer() *gin.Engine {
	if !pkg.IsDevEnv() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	//设置上传文件大小限制
	r.MaxMultipartMemory = app.Config.App.UploadMaxM << 20
	r.Use(gin.Recovery())
	//跨域中间件
	if app.Config.Http.UseCrossMiddleware {
		r.Use(middleware.CrossMiddleware())
	}
	r.Use(middleware.UUID())

	return r
}
