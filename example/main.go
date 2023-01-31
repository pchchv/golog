package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pchchv/golog"
	"github.com/pkg/errors"
)

func someFunc() error {
	// This function simulated a library/framework throwing an error but not using the "errors" package. Therefor this
	// pure go error doesn't contain any stack trace information.
	return fmt.Errorf("BOOM!!! some error occurred maybe from within an framework?")
}

func thatFunc() error {
	// The errors package will add stack trace information which is later used by golog to print that stack trace
	return errors.Wrap(someFunc(), "that func wrapped this error")
}

func thisFunc() error {
	return thatFunc()
}

func main() {
	golog.Plain("Hello world!") // will not be printed as log level is to restrictive
	golog.Info("Hello world!")
	golog.Debug("Hello world %d times!", 42) // not shown because log-level is on INFO
	golog.Error("Hello world!")

	golog.LogLevel = golog.LOG_PLAIN
	golog.Plain("Some plain text") // now the log level is ok

	time.Sleep(time.Millisecond)
	fmt.Print("\n===== 1 =====\n")
	golog.LogLevel = golog.LOG_DEBUG

	golog.Info("Hello %s!", "world")
	golog.Debug("Hello world %d times!", 42) // shown because log-level is on DEBUG
	golog.Error("Hello %x!", 123)

	time.Sleep(time.Millisecond)
	fmt.Print("\n===== 2 =====\n")
	golog.FormatFunctions[golog.LOG_INFO] = simpleInfo

	golog.Info("Some")
	golog.Info("AMAZING")
	golog.Info("log")
	golog.Info("entries")
	golog.Debug("Boring")
	golog.Error("Lame")

	time.Sleep(time.Millisecond)
	fmt.Print("\n===== 3 =====\n")
	// This will print stack trace information because "thatFunc()" added them:
	golog.Stack(thisFunc())
	// This doesn't and just prints the error message:
	golog.Stack(someFunc())

	time.Sleep(time.Millisecond)
	fmt.Print("\n===== 4 =====\n")
	golog.DateFormat = "02.01.2006 at 15:04:05"

	golog.Info("Hello world!")
	golog.Debug("Hello world!")
	golog.Error("Hello world!")

	golog.FatalCheck(thisFunc())
	golog.Panic("Hello world!")
}

func simpleInfo(writer *os.File, time, level string, maxLength int, caller, message string) {
	fmt.Fprintf(writer, ">>  My custom Info  ||  %s\n", message)
}
