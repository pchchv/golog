package golog

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	Text string
	Err  error
	Time time.Time
	File *os.File
}

func NewLogger(f *os.File) *Logger {
	l := &Logger{}
	l.File = f
	return l
}

func (l *Logger) Print() {
	fmt.Printf("%v: %s", l.Time, l.Text)
}

func (l *Logger) Log() {
	l.File.WriteString(fmt.Sprintf("%v: %s", l.Time, l.Text))
}

func (l *Logger) Panic() {
	panic(l.Err)
}

func (l *Logger) Error() {
	l.File.WriteString(fmt.Sprintf("%v: ERROR: %v", l.Time, l.Err))
	fmt.Println(l.Err.Error())
}
