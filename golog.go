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

func Plain(format string, a ...interface{}) {
	if LogLevel <= LOG_PLAIN {
		log(LOG_PLAIN, 3, fmt.Sprintf(format, a...))
	}
}

// Plainb is equal to Plain(...) but can go back in the stack and can therefore show function positions from previous functions.
func Plainb(framesBackward int, format string, a ...interface{}) {
	if LogLevel <= LOG_PLAIN {
		log(LOG_PLAIN, 3+framesBackward, fmt.Sprintf(format, a...))
	}
}

func Info(format string, a ...interface{}) {
	if LogLevel <= LOG_INFO {
		log(LOG_INFO, 3, fmt.Sprintf(format, a...))
	}
}

// Infob is equal to Info(...) but can go back in the stack and can therefore show function positions from previous functions.
func Infob(framesBackward int, format string, a ...interface{}) {
	if LogLevel <= LOG_INFO {
		log(LOG_INFO, 3+framesBackward, fmt.Sprintf(format, a...))
	}
}

func Debug(format string, a ...interface{}) {
	if LogLevel <= LOG_DEBUG {
		log(LOG_DEBUG, 3, fmt.Sprintf(format, a...))
	}
}

// Debugb is equal to Debug(...) but can go back in the stack and can therefore show function positions from previous functions.
func Debugb(framesBackward int, format string, a ...interface{}) {
	if LogLevel <= LOG_DEBUG {
		log(LOG_DEBUG, 3+framesBackward, fmt.Sprintf(format, a...))
	}
}

func Error(format string, a ...interface{}) {
	if LogLevel <= LOG_ERROR {
		log(LOG_ERROR, 3, fmt.Sprintf(format, a...))
	}
}

// Errorb is equal to Error(...) but can go back in the stack and can therefore show function positions from previous functions.
func Errorb(framesBackward int, format string, a ...interface{}) {
	if LogLevel <= LOG_ERROR {
		log(LOG_ERROR, 3+framesBackward, fmt.Sprintf(format, a...))
	}
}

// Stack tries to print the stack trace of the given error using the  %+v  format string.
// When using the https://github.com/pkg/errors package, a full error stack trace will be output.
// If normal errors are used, just print the error.
func Stack(err error) {
	if LogLevel <= LOG_ERROR {
		// Directly call "log" to avoid extra function call
		log(LOG_ERROR, 3, fmt.Sprintf("%+v", err))
	}
}

// Stackb is equal to Stack(...) but can go back in the stack and can therefore show function positions from previous functions.
func Stackb(framesBackward int, err error) {
	if LogLevel <= LOG_ERROR {
		// Directly call "log" to avoid extra function call
		log(LOG_ERROR, 3+framesBackward, fmt.Sprintf("%+v", err))
	}
}

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
