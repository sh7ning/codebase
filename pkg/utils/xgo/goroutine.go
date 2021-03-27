package xgo

import (
	"codebase/pkg/log"
	"fmt"
	"runtime"
	"time"

	"go.uber.org/zap"
)

// Go goroutine
func Go(fn func()) {
	Try(fn, nil)
}

// Delay goroutine
func Delay(delay time.Duration, fn func()) {
	Try(func() {
		time.Sleep(delay)
		fn()
	}, nil)
}

func Try(fn func(), cleaner func(error)) {
	go func() {
		defer func() {
			var err error
			if r := recover(); r != nil {
				_, file, line, _ := runtime.Caller(2)
				log.Error("recover", zap.Any("err", r), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
				if e, ok := r.(error); ok {
					err = e
				} else {
					err = fmt.Errorf("%+v", r)
				}
			}
			if cleaner != nil {
				defer func() {
					if r := recover(); r != nil {
						_, file, line, _ := runtime.Caller(2)
						log.Error("recover cleaner", zap.Any("err", r), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
					}
				}()
				cleaner(err)
			}
		}()
		fn()
	}()

	return
}
