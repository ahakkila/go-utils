// Package logger provides functionality for logging on various loglevels
package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

type Logger struct {
	countLines  bool
	countChars  bool
	fatalWriter io.Writer
	errorWriter io.Writer
	warnWriter  io.Writer
	infoWriter  io.Writer
	traceWriter io.Writer
}

// Create a new logger.Logger
func NewLogger() *Logger {
	return &Logger{
		countLines:  false,
		countChars:  false,
		fatalWriter: ioutil.Discard,
		errorWriter: ioutil.Discard,
		warnWriter:  ioutil.Discard,
		infoWriter:  ioutil.Discard,
		traceWriter: ioutil.Discard,
	}
}

// Redirect all output to stdout
func (l *Logger) AllStdout() *Logger {
	l.fatalWriter = os.Stdout
	l.errorWriter = os.Stdout
	l.warnWriter = os.Stdout
	l.infoWriter = os.Stdout
	l.traceWriter = os.Stdout
	return l
}

// Redirect Fatal loglevel to stdout
func (l *Logger) FatalStdout() *Logger {
	return l.FatalOutput(os.Stdout)
}

// Redirect Error loglevel to stdout
func (l *Logger) ErrorStdout() *Logger {
	return l.ErrorOutput(os.Stdout)
}

// Redirect Warning loglevel to stdout
func (l *Logger) WarnStdout() *Logger {
	return l.WarnOutput(os.Stdout)
}

// Redirect Information loglevel to stdout
func (l *Logger) InfoStdout() *Logger {
	return l.InfoOutput(os.Stdout)
}

// Redirect Trace loglevel to stdout
func (l *Logger) TraceStdout() *Logger {
	return l.TraceOutput(os.Stdout)
}

// Discard all output, redirected to ioutil.Discard
func (l *Logger) DiscardAll() *Logger {
	l.fatalWriter = ioutil.Discard
	l.errorWriter = ioutil.Discard
	l.warnWriter = ioutil.Discard
	l.infoWriter = ioutil.Discard
	l.traceWriter = ioutil.Discard
	return l
}

// Discard Fatal output, redirected to ioutil.Discard
func (l *Logger) FatalDiscard() *Logger {
	return l.FatalOutput(ioutil.Discard)
}

// Discard Error output, redirected to ioutil.Discard
func (l *Logger) ErrorDiscard() *Logger {
	return l.ErrorOutput(ioutil.Discard)
}

// Discard Warning output, redirected to ioutil.Discard
func (l *Logger) WarnDiscard() *Logger {
	return l.WarnOutput(ioutil.Discard)
}

// Discard Information output, redirected to ioutil.Discard
func (l *Logger) InfoDiscard() *Logger {
	return l.InfoOutput(ioutil.Discard)
}

// Discard Trace output, redirected to ioutil.Discard
func (l *Logger) TraceDiscard() *Logger {
	return l.TraceOutput(ioutil.Discard)
}

// Redirect all output to w
func (l *Logger) AllOutput(w io.Writer) *Logger {
	l.fatalWriter = w
	l.errorWriter = w
	l.warnWriter = w
	l.infoWriter = w
	l.traceWriter = w
	return l
}

// Redirect Fatal output to w
func (l *Logger) FatalOutput(w io.Writer) *Logger {
	l.fatalWriter = w
	return l
}

// Redirect Error output to w
func (l *Logger) ErrorOutput(w io.Writer) *Logger {
	l.errorWriter = w
	return l
}

// Redirect Warning output to w
func (l *Logger) WarnOutput(w io.Writer) *Logger {
	l.warnWriter = w
	return l
}

// Redirect Information output to w
func (l *Logger) InfoOutput(w io.Writer) *Logger {
	l.infoWriter = w
	return l
}

// Redirect Trace output to w
func (l *Logger) TraceOutput(w io.Writer) *Logger {
	l.traceWriter = w
	return l
}

// Log a message on Fatal loglevel and call panic()
func (l *Logger) Fatal(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	panicline := fmt.Sprintf(" [%s] %s (fatal): %s", now(), f, line)
	fmt.Fprintln(l.fatalWriter, panicline)
	panic(panicline)
}

// Log a message on Error loglevel
func (l *Logger) Error(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.errorWriter, " [%s] %s (error): %s\n", now(), f, line)
}

// Log a message on Warning loglevel
func (l *Logger) Warn(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.warnWriter, " [%s] %s (warn): %s\n", now(), f, line)
}

// Log a message on Information loglevel
func (l *Logger) Info(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.infoWriter, " [%s] %s (info): %s\n", now(), f, line)
}

// Log a message on Trace loglevel
func (l *Logger) Trace(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.traceWriter, " [%s] %s (trace): %s\n", now(), f, line)
}

// Return timestring in "15:04:05" format
func now() string {
	return time.Now().Format("15:04:05")
}

// Return name of the calling function
func FName() string {
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	return f.Name()
}

// One level further in the call stack as this is called by Logger
func FNameForLogger() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	return f.Name()
}
