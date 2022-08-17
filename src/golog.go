package golog

import "sync"

// Golog is base struct
type Golog struct {
	logger     loggers
	level      uint32
	buffer     sync.Pool
	enableJSON bool
}
