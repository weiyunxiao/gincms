package router

import (
	"gincms/app"
	"gincms/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.GET("ping", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Use(middleware.UUID())
	if app.Config.App.IsDemo {
		r.Use(middleware.CheckDemoEnv())
	}
	AdminApiRouter(r)
}
