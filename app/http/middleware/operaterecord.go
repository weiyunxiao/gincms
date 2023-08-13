package middleware

import (
	"fmt"
	"gincms/app"
	"gincms/app/http/adminapi/service/sys"
	"gincms/app/model"
	"gincms/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// OperationRecord 记录用户的操作(不记录get操作)
func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			return
		}
		record := model.SysLogOperate{
			ID:         0,
			ReqUri:     c.Request.URL.Path,
			ReqMethod:  c.Request.Method,
			Ip:         c.ClientIP(),
			UserAgent:  c.Request.UserAgent(),
			Status:     0,
			ResultMsg:  "",
			TenantID:   0,
			CreateTime: carbon.DateTime{},
		}
		now := time.Now()
		c.Next()
		record.UserID = c.GetInt64("uid")
		record.Duration = fmt.Sprintf("%v", time.Since(now))
		record.CreateTime = carbon.DateTime{Carbon: carbon.Now()}
		err := sys.LogService.AddOperateLog(c, &record)
		if err != nil {
			app.Logger.Error("记录操作日志错误", zap.String("reqKey", pkg.GetReqKey(c)), zap.Error(err))
		}
	}
}
