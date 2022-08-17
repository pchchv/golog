package golog

import (
	"io"
	"sync"
)

// Golog is base struct
type Golog struct {
	logger     loggers
	level      uint32
	buffer     sync.Pool
	enableJSON bool
}

// MODE is logging mode (std only, writer only, std & writer)
type MODE uint8

type wMode uint8

type traceMode int64

type logger struct {
	tag              string
	rawtag           []byte
	writer           io.Writer
	std              io.Writer
	color            func(string) string
	isColor          bool
	traceMode        traceMode
	mode             MODE
	prevMode         MODE
	writeMode        wMode
	disableTimestamp bool
}
