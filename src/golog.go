package golog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Golog is base struct
type Golog struct {
	L              Logger
	PrintToFile    bool // If false doesn't write to the file
	PrintToConsole bool // If false doesn't write to the console
	File           *os.File
}

func New(print bool, saving bool, path string) (*Golog, error) {
	var err error
	g := &Golog{}
	g.PrintToConsole = print
	g.PrintToFile, g.File, err = enablingFile(saving, path)
	if err != nil {
		return g, err
	}
	defer g.File.Close()
	g.L = *NewLogger()
	return g, nil
}

func enablingFile(s bool, p string) (bool, *os.File, error) {
	var f *os.File
	e := errors.New("The file path was set incorrectly")
	if s {
		if p == "" {
			return s, f, e
		}
		p, err := filepath.Abs(p)
		if err != nil {
			return s, f, e
		}
		f, err = os.Create(p)
		if err != nil {
			return s, f, errors.New(fmt.Sprintf("File creation error: %v", err))
		}
	}
	return s, f, nil
}

func (g *Golog) Print(text string) {
	g.L.Time = time.Now()
	g.L.Text = text
	g.L.Print()
}

func (g *Golog) Log(text string) {
	g.L.Time = time.Now()
	g.L.Text = text
	if g.PrintToConsole {
		g.L.Print()
	}
	if g.PrintToFile {
		g.L.Log(g.File)
	}
}

func (g *Golog) Panic(err error) {
	g.L.Time = time.Now()
	g.L.Error = err
	g.L.Text = fmt.Sprintf("Panic: %v", err)
	if g.PrintToFile {
		g.L.Log(g.File)
	}
	g.L.Panic()
}
