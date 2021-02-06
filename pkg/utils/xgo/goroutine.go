package xgo

import (
	"codebase/pkg/log"
	"fmt"
	"runtime"
	"time"

	"go.uber.org/zap"
)

// Go goroutine
func Go(fn func()) (ec chan error) {
	return Try(fn, nil)
}

// Delay goroutine
func Delay(delay time.Duration, fn func()) {
	Try(func() {
		time.Sleep(delay)
		fn()
	}, nil)
}

func Try(fn func(), cleaner func()) (ec chan error) {
	ec = make(chan error, 1)
	go func() {
		defer func() {
			defer close(ec)
			if err := recover(); err != nil {
				_, file, line, _ := runtime.Caller(2)
				log.Error("recover", zap.Any("err", err), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
				if e, ok := err.(error); ok {
					ec <- e
				} else {
					ec <- fmt.Errorf("%+v", err)
				}

				if cleaner != nil {
					defer func() {
						if err := recover(); err != nil {
							_, file, line, _ := runtime.Caller(2)
							log.Error("recover cleaner", zap.Any("err", err), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
						}
					}()
					cleaner()
				}
			}
		}()
		fn()
	}()

	return
}
