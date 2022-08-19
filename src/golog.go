package golog

import (
	"errors"
	"path/filepath"
	"time"
)

// Golog is base struct
type Golog struct {
	L              Logger
	PrintToFile    bool // If false doesn't write to the file
	PrintToConsole bool // If false doesn't write to the console
	FilePath       string
}

func New(print bool, saving bool, path string) (*Golog, error) {
	var err error
	g := &Golog{}
	g.PrintToConsole = print
	g.PrintToFile, g.FilePath, err = enablingFile(saving, path)
	if err != nil {
		return g, err
	}
	g.L = *NewLogger()
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
	g.L.Time = time.Now()
	g.L.Text = text
	g.L.Print()
}

func (g *Golog) Log(text string) {
	g.L.Text = text
	if g.PrintToConsole {
		g.L.Print()
	}
	if g.PrintToFile {
	}
}
