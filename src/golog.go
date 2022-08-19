package golog

import (
	"errors"
	"path/filepath"
	"time"
)

// Golog is base struct
type Golog struct {
	l      Logger
	saving bool // If false doesn't write to the file
	print  bool // If false doesn't write to the console
	path   string
}

func New(print bool, saving bool, path string) (*Golog, error) {
	var err error
	g := &Golog{}
	g.print = print
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
	// Need to implement the creation of a file
	return s, p, nil
}

func (g *Golog) Print(text string) {
	g.l.time = time.Now()
	g.l.text = text
	g.l.Print()
}

func (g *Golog) Log(text string) {
	g.l.text = text
	if g.print {
		g.l.Print()
	}
	if g.saving {
	}
}
