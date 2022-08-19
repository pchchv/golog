package golog

import (
	"fmt"
	"time"
)

type Logger struct {
	Text  string
	Error error
	Time  time.Time
}

func NewLogger() *Logger {
	l := &Logger{}
	return l
}

func (l *Logger) Print() {
	fmt.Printf("%v: %s", l.Time, l.Text)
}

func (l *Logger) Log(g *Golog) {
	// need implement a print to file
}
