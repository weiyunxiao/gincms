package http

import (
	"gincms/app"
	"gincms/app/http/middleware"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
)

// gin初始化
func CreateGinServer() *gin.Engine {
	if !pkg.IsDevEnv() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	//跨域中间件
	if app.Config.Http.UseCrossMiddleware {
		r.Use(middleware.CrossMiddleware())
	}
	r.Use(middleware.UUID())

	return r
}
