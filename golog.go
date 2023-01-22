package golog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type Level int

const (
	LOG_PLAIN Level = iota
	LOG_DEBUG
	LOG_INFO
	LOG_ERROR
	LOG_FATAL
)

var (
	LogLevel   = LOG_INFO
	DateFormat = "2006-01-02 15:04:05.000"

	FormatFunctions = map[Level]func(*os.File, string, string, int, string, string){
		LOG_PLAIN: LogPlain,
		LOG_DEBUG: LogDefault,
		LOG_INFO:  LogDefault,
		LOG_ERROR: LogDefault,
		LOG_FATAL: LogDefault,
	}

	// The current maximum length printed for caller information. This is updated each time something gets printed
	CallerColumnWidth = 0

	LevelStrings = map[Level]string{
		LOG_PLAIN: "",
		LOG_DEBUG: "[DEBUG]",
		LOG_INFO:  "[INFO] ",
		LOG_ERROR: "[ERROR]",
		LOG_FATAL: "[FATAL]",
	}

	LevelOutputs = map[Level]*os.File{
		LOG_PLAIN: os.Stdout,
		LOG_DEBUG: os.Stdout,
		LOG_INFO:  os.Stdout,
		LOG_ERROR: os.Stderr,
		LOG_FATAL: os.Stderr,
	}
)

func log(level Level, framesBackward int, message string) {
	// We know here that the stack contains two calls from inside this file.
	// The third frame comes from the file that initially called a function
	// in this file e.g. Info()
	caller := getCallerDetails(framesBackward)

	updateCallerColumnWidth(caller)

	FormatFunctions[level](LevelOutputs[level], time.Now().Format(DateFormat), LevelStrings[level], CallerColumnWidth, caller, message)
}

func getCallerDetails(framesBackwards int) string {
	name := ""
	line := -1
	ok := false

	if _, name, line, ok = runtime.Caller(framesBackwards); ok {
		name = path.Base(name)
	}

	caller := fmt.Sprintf("%s:%d", name, line)

	return caller
}

func updateCallerColumnWidth(caller string) {
	if len(caller) > CallerColumnWidth {
		CallerColumnWidth = len(caller)
	}
}

func LogPlain(writer *os.File, time, level string, maxLength int, caller, message string) {
	fmt.Fprintf(writer, "%s\n", message)
}

func LogDefault(writer *os.File, time, level string, maxLength int, caller, message string) {
	fmt.Fprintf(writer, "%s %s %-*s | %s\n", time, level, maxLength, caller, message)
}
