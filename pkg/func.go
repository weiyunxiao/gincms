package pkg

import (
	"gincms/app"
	"gincms/pkg/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

// IsDevEnv 判断是否是开发环境
func IsDevEnv() bool {
	return strings.HasPrefix(app.Config.App.Env, "dev")
}

// ParamErrLog 记录参数绑定错误
func ParamErrLog(c *gin.Context, err error) {
	if app.Config.App.LogParamErr {
		app.Logger.Warn("参数错误", zap.String("reqKey", GetReqKey(c)), zap.Error(err))
	}
}

// SortStr 构建排序字符串
func SortStr(orderField string, sortAsc bool, defaultSortField ...string) (orderStr string) {
	sortType := " desc"
	if sortAsc {
		sortType = " asc"
	}
	if len(orderField) > 0 {
		orderField = util.FilterSQLInjection(orderField)
		return orderField + sortType
	}

	if len(defaultSortField) == 0 {
		return "sort " + sortType
	}
	return defaultSortField[0] + sortType
}

// PaginateScope gorm分页Scope
func PaginateScope(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case limit > 2000:
			limit = 2000
		case limit <= 0:
			limit = 10
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

// GetReqKey 获取reqKey
func GetReqKey(c *gin.Context) string {
	return c.GetString("reqKey")
}

func GetContentKey(c *gin.Context, key string) string {
	return c.GetString("reqKey")
}
