package bootstrap

import (
	"flag"
	"fmt"
	"gincms/app"
	"gincms/app/bootstrap/internal"
	"gincms/pkg"
	"gincms/pkg/util"
	"github.com/gookit/goutil/fsutil"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
)

// AppInit 一些初始化的操作
func AppInit() {
	InitConfig()     //加载应用的各种配置信息
	InitPathInfo()   //项目路径相关定义
	initZap()        //日志组件
	InitDB()         //初始化DB列表
	InitLocalCache() //初始化本地缓存
}

// InitConfig 使用viper处理应用目录下的config.yaml
func InitConfig() {
	var configFile = flag.String("f", "config.yaml", "the config file")
	flag.Parse()

	viper.SetConfigFile(*configFile)             // 指定配置文件
	viper.AddConfigPath(".")                     // 去哪里找配置文件，这里'.'表示main.go同级目录下
	if err := viper.ReadInConfig(); err != nil { // 查找并读取配置文件
		log.Fatalln("加载应用配置文件时出错:", err)
	}

	if err := viper.Unmarshal(&app.Config); err != nil {
		log.Fatalln("解析配置文件到config时出错:", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&app.Config); err != nil {
			log.Println("配置文件被修改,重新解析配置文件到config时出错:", err)
		} else {
			log.Println("重新加载配置文件")
		}
	})
}

// InitPathInfo 设置相关目录变量
func InitPathInfo() {
	app.Config.PathInfo.Root = filepath.Dir(filepath.Dir(util.CurDir()))
}

// InitDB 数据库的初始化
func InitDB() {
	app.DBList = map[string]*gorm.DB{}
	var loggerMode logger.Interface
	if pkg.IsDevEnv() {
		loggerMode = logger.Default.LogMode(logger.Info)
	} else {
		loggerMode = logger.Default.LogMode(logger.Silent)
	}
	for key, dbStruct := range app.Config.DBDefineList {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dbStruct.Dsn, // DSN data source name
			DefaultStringSize:         191,          // string 类型字段的默认长度
			DisableDatetimePrecision:  true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,        // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
			Logger: loggerMode,
		})
		if err != nil {
			log.Fatalf("数据库错误-(‘%s’定义):%v\r\n", key, err)
		}
		TheDB, _ := db.DB()
		TheDB.SetMaxIdleConns(dbStruct.MaxIdleConns)
		TheDB.SetMaxOpenConns(dbStruct.MaxOpenConns)
		TheDB.SetConnMaxLifetime(time.Duration(dbStruct.ConnMaxLifetime) * time.Second)
		TheDB.SetConnMaxIdleTime(time.Duration(dbStruct.ConnMaxIdleTime) * time.Second)
		app.DBList[key] = db
	}
}

// InitLocalCache 本地缓存的初始化
func InitLocalCache() {
	app.LocalCache = cache.New(5*time.Minute, 10*time.Minute)
}

// 日志组件的初始化
func initZap() {
	logPath := app.Config.App.LogPath
	if !fsutil.PathExist(logPath) { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", logPath)
		_ = os.Mkdir(logPath, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	loggerObj := zap.New(zapcore.NewTee(cores...))
	app.Logger = loggerObj.WithOptions(zap.AddCaller())
}
