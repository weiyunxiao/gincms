package router

import (
	"gincms/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.GET("ping", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Use(middleware.UUID())
	AdminApiRouter(r)
}
