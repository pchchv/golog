package golog

import (
	"fmt"
	"os"
)

func LogPlain(writer *os.File, time, level string, maxLength int, caller, message string) {
	fmt.Fprintf(writer, "%s\n", message)
}

func LogDefault(writer *os.File, time, level string, maxLength int, caller, message string) {
	fmt.Fprintf(writer, "%s %s %-*s | %s\n", time, level, maxLength, caller, message)
}
