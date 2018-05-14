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
	CountLines  bool
	CountChars  bool
	FatalWriter io.Writer
	ErrorWriter io.Writer
	WarnWriter  io.Writer
	InfoWriter  io.Writer
	TraceWriter io.Writer
}

func NewLogger() *Logger {
	return &Logger{
		CountLines:  false,
		CountChars:  false,
		FatalWriter: ioutil.Discard,
		ErrorWriter: ioutil.Discard,
		WarnWriter:  ioutil.Discard,
		InfoWriter:  ioutil.Discard,
		TraceWriter: ioutil.Discard,
	}
}

func (l *Logger) AllStdout() *Logger {
	l.FatalWriter = os.Stdout
	l.ErrorWriter = os.Stdout
	l.WarnWriter = os.Stdout
	l.InfoWriter = os.Stdout
	l.TraceWriter = os.Stdout
	return l
}

func (l *Logger) FatalStdout() *Logger {
	return l.FatalOutput(os.Stdout)
}

func (l *Logger) ErrorStdout() *Logger {
	return l.ErrorOutput(os.Stdout)
}

func (l *Logger) WarnStdout() *Logger {
	return l.WarnOutput(os.Stdout)
}

func (l *Logger) InfoStdout() *Logger {
	return l.InfoOutput(os.Stdout)
}

func (l *Logger) TraceStdout() *Logger {
	return l.TraceOutput(os.Stdout)
}

func (l *Logger) DiscardAll() *Logger {
	l.FatalWriter = ioutil.Discard
	l.ErrorWriter = ioutil.Discard
	l.WarnWriter = ioutil.Discard
	l.InfoWriter = ioutil.Discard
	l.TraceWriter = ioutil.Discard
	return l
}

func (l *Logger) FatalDiscard() *Logger {
	return l.FatalOutput(ioutil.Discard)
}

func (l *Logger) ErrorDiscard() *Logger {
	return l.ErrorOutput(ioutil.Discard)
}

func (l *Logger) WarnDiscard() *Logger {
	return l.WarnOutput(ioutil.Discard)
}

func (l *Logger) InfoDiscard() *Logger {
	return l.InfoOutput(ioutil.Discard)
}

func (l *Logger) TraceDiscard() *Logger {
	return l.TraceOutput(ioutil.Discard)
}

func (l *Logger) AllOutput(w io.Writer) *Logger {
	l.FatalWriter = w
	l.ErrorWriter = w
	l.WarnWriter = w
	l.InfoWriter = w
	l.TraceWriter = w
	return l
}

func (l *Logger) FatalOutput(w io.Writer) *Logger {
	l.FatalWriter = w
	return l
}

func (l *Logger) ErrorOutput(w io.Writer) *Logger {
	l.ErrorWriter = w
	return l
}

func (l *Logger) WarnOutput(w io.Writer) *Logger {
	l.WarnWriter = w
	return l
}

func (l *Logger) InfoOutput(w io.Writer) *Logger {
	l.InfoWriter = w
	return l
}

func (l *Logger) TraceOutput(w io.Writer) *Logger {
	l.TraceWriter = w
	return l
}

func (l *Logger) Fatal(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	panicline := fmt.Sprintf(" [%s] %s (fatal): %s", l.Now(), f, line)
	fmt.Fprintln(l.FatalWriter, panicline)
	panic(panicline)
}

func (l *Logger) Error(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.ErrorWriter, " [%s] %s (error): %s\n", l.Now(), f, line)
}

func (l *Logger) Warn(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.WarnWriter, " [%s] %s (warn): %s\n", l.Now(), f, line)
}

func (l *Logger) Info(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.InfoWriter, " [%s] %s (info): %s\n", l.Now(), f, line)
}

func (l *Logger) Trace(format string, a ...interface{}) {
	f := FNameForLogger()
	line := fmt.Sprintf(format, a...)
	fmt.Fprintf(l.TraceWriter, " [%s] %s (trace): %s\n", l.Now(), f, line)
}

func (l *Logger) Now() string {
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
