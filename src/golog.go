package golog

// Golog is base struct
type Golog struct {
	l      logger
	saving bool // If false doesn't write to the file
	file   string
}
