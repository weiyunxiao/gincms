package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqKeyFunc := uuid.New().String
		c.Set("reqKey", reqKeyFunc())
		c.Next()
	}
}
