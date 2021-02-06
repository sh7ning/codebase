package db

import (
	"codebase/pkg/log"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	gl "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type logger struct {
}

func (l *logger) LogMode(level gl.LogLevel) gl.Interface {
	return l
}

func (l *logger) Info(_ context.Context, s string, args ...interface{}) {
	log.Info(fmt.Sprintf(s, args...))
}

func (l *logger) Warn(_ context.Context, s string, args ...interface{}) {
	log.Warn(fmt.Sprintf(s, args...))
}

func (l *logger) Error(_ context.Context, s string, args ...interface{}) {
	log.Error(fmt.Sprintf(s, args...))
}

func (l *logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	log.Debug(utils.FileWithLineNum(), zap.String("sql", sql), zap.Int64("rows", rows), zap.Duration("elapsed", elapsed), zap.Error(err))
}
