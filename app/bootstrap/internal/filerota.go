package internal

import (
	"gincms/app"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(app.Config.App.LogPath, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(7)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	return zapcore.AddSync(fileWriter), err
}
