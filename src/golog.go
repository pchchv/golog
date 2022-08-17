package golog

import (
	"bytes"
	"io"
	"sync"
	"sync/atomic"
)

// Golog is base struct
type Golog struct {
	bs          *uint64
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

// LEVEL is log level
type LEVEL uint8

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

const (
	// DEBG is debug log level
	DEBG LEVEL = iota + 1
	// FATAL is fatal log level
	FATAL
	timeFormat         = "2006-01-02 15:04:05"
	tab                = "\t"
	lsep               = tab + "["
	lsepl              = len(lsep)
	sep                = "]:" + tab
	sepl               = len(sep)
	DefaultCallerDepth = 2
)

// New returns plain glg instance
func New() *Golog {
	g := &Golog{
		level:       new(uint32),
		callerDepth: DefaultCallerDepth,
	}
	g.bs = new(uint64)

	atomic.StoreUint64(g.bs, uint64(len(timeFormat)+lsepl+sepl))

	g.buffer = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, int(atomic.LoadUint64(g.bs))))
		},
	}

	atomic.StoreUint32(g.level, uint32(FATAL))

	return g
}
