package golog

import "time"

type logger struct {
	text string
	err  error
	time time.Time
}
