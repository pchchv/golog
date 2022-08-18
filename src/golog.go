package golog

import (
	"errors"
	"path/filepath"
)

// Golog is base struct
type Golog struct {
	l      logger
	saving bool // If false doesn't write to the file
	path   string
}

func New(saving bool, path string) (*Golog, error) {
	var err error
	g := &Golog{}
	g.saving, g.path, err = enablingFile(saving, path)
	if err != nil {
		return g, err
	}
	g.l = *NewLogger()
	return g, nil
}

func enablingFile(s bool, p string) (bool, string, error) {
	e := errors.New("The file path was set incorrectly")
	if s {
		if p == "" {
			return s, p, e
		}
		p, err := filepath.Abs(p)
		if err != nil {
			return s, p, e
		}
	}
	return s, p, nil
}
