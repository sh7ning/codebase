//package log
//see: https://github.com/uber-go/zap
//demo:
//	_ = log.New(false)
//	defer log.Sync()
//	log.Info("this is a test message")
package log

import (
	"app/pkg/cfg"
	"app/pkg/services/dingtalk"
	"app/pkg/utils/helper"
	"app/pkg/utils/recovery"
	"errors"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func Logger() *zap.Logger {
	return logger
}

func SetLogger(log *zap.Logger) error {
	if logger != nil {
		return errors.New("logger is not nil")
	}

	logger = log
	return nil
}

func New(development bool) {
	if logger != nil {
		return
	}

	var config Config

	if development {
		config = NewDevelopment()
	} else {
		config = NewProduction()
	}

	opts := make([]zap.Option, 0, 1)
	if cfg.AppConfig.App.ErrorReporting {
		//warn 级别以上发送钉钉消息
		opts = append(opts, zap.Hooks(func(entry zapcore.Entry) error {
			if zap.WarnLevel.Enabled(entry.Level) {
				go func() {
					defer recovery.Recover(func(stack string) {
						fmt.Println("logger recovery, 钉钉发送消息失败:" + stack)
					})()

					//!!! 此处会修改 entry.Stack 防止 stacktrace 过长
					stack := []rune(entry.Stack)
					if len(stack) > 2048 {
						entry.Stack = string(stack[:2048])
					}
					ding := dingtalk.Instance("")
					if ding != nil {
						if err := ding.SendTextMessage(helper.ToJsonString(entry), nil, false); err != nil {
							fmt.Println("钉钉发送消息失败:" + err.Error())
						}
					}
				}()
			}

			return nil
		}))
	}

	logger = config.Build(opts...)
}

func NewDevelopment() Config {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	w := os.Stdout
	sink := zapcore.AddSync(w)

	return Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoder:     zapcore.NewConsoleEncoder(encoderConfig),
		WriteSyncer: sink,
	}
}

func NewProduction() Config {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//https://github.com/natefinch/lumberjack
	var w io.Writer
	w = os.Stdout
	if cfg.AppConfig.App.LogFile != "" {
		w = &lumberjack.Logger{
			Filename:   cfg.AppConfig.App.LogFile,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     1,    //days
			Compress:   true, // disabled by default
		}
	}
	sink := zapcore.AddSync(w)

	return Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoder:     zapcore.NewJSONEncoder(encoderConfig),
		WriteSyncer: sink,
		InitialFields: map[string]interface{}{
			"app_name": cfg.AppConfig.App.Name,
		},
		//Sampling: &zap.SamplingConfig{
		//	Initial:    100,
		//	Thereafter: 100,
		//},
	}
}

func Sync() error {
	return logger.Sync()
}

//
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

//DPanic DPanic means "development panic"
//Deprecated: 不建议采用
func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
