package golog

import (
	"io"
	"sync"
)

// Golog is base struct
type Golog struct {
	logger      loggers
	level       *uint32
	buffer      sync.Pool
	callerDepth int
	enableJSON  bool
}

// JSONFormat is json object structure for logging
type JSONFormat struct {
	Date   string      `json:"date,omitempty"`
	Level  string      `json:"level,omitempty"`
	File   string      `json:"file,omitempty"`
	Detail interface{} `json:"detail,omitempty"`
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
