package golog

import (
	"sync"
	"sync/atomic"
)

type loggers struct {
	mu     sync.Mutex
	read   atomic.Value
	misses int
}
