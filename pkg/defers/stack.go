package defers

import (
	"sync"
)

func newStack() *deferStack {
	return &deferStack{
		fns: make([]func() error, 0),
		mu:  sync.RWMutex{},
	}
}

type deferStack struct {
	fns []func() error
	mu  sync.RWMutex
}

func (ds *deferStack) push(fns ...func() error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.fns = append(ds.fns, fns...)
}

func (ds *deferStack) clean() {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	for i := len(ds.fns) - 1; i >= 0; i-- {
		_ = ds.fns[i]()
	}
}
