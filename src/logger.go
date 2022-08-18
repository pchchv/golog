package golog

import "time"

type logger struct {
	text string
	err  error
	time time.Time
}

func NewLogger() *logger {
	l := &logger{}
	return l
}
