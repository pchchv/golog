package golog

import (
	"fmt"
	"time"
)

type Logger struct {
	text string
	err  error
	time time.Time
}

func NewLogger() *Logger {
	l := &Logger{}
	return l
}

func (l *Logger) Print() {
	fmt.Printf("%v: %s", l.time, l.text)
}

func (l *Logger) Log(g *Golog) {
	// need implement a print to file
}
