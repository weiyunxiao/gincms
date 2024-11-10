package app

import (
	"gincms/config"
	"strings"

	"github.com/go-redsync/redsync/v4"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config      config.Config
	Logger      *zap.Logger
	LocalCache  *cache.Cache
	DBList      map[string]*gorm.DB
	Redis       *redis.Client
	RedisLocker *redsync.Redsync //用法参考https://github.com/go-redsync/redsync
)

// DB 获取定义的DB实例
func DB(dbKeyArr ...string) *gorm.DB {
	dbKey := "Default"
	if len(dbKeyArr) > 0 {
		dbKey = dbKeyArr[0]
		dbKey = strings.ToLower(dbKey) //map解析时是小写字母
		if _, ok := DBList[dbKey]; ok {
			return DBList[dbKey]
		}
	}

	dbKey = strings.ToLower(dbKey) //map解析时是小写字母
	return DBList[dbKey]
}
