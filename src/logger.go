package golog

import (
	"fmt"
	"os"
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

func (l *Logger) Log(f *os.File) {
	f.WriteString(fmt.Sprintf("%v: %s", l.Time, l.Text))
}

func (l *Logger) Panic() {
	panic(l.Error)
}
