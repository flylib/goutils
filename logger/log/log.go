package log

import (
	"fmt"
	"log"
	"os"
)

// Level These are the integer logging levels used by the logger
type Level int

// Comment
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	levelFlags = []string{"DEBG", "INFO", "WARN", "ERRO", "FATL"}
)

type logger struct {
	*log.Logger
	callDepth int
}

func NewLogger(options ...Option) *logger {
	newLogger := new(logger)
	newLogger.callDepth = 2
	for _, f := range options {
		f(newLogger)
	}
	return newLogger
}

func (l *logger) Info(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[INFO], fmt.Sprint(args...)))
}

func (l *logger) Warn(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[WARNING], fmt.Sprint(args...)))
}

func (l *logger) Debug(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[DEBUG], fmt.Sprint(args...)))
}

func (l *logger) Error(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[ERROR], fmt.Sprint(args...)))
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *logger) Fatal(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[FATAL], fmt.Sprint(args...)))
	os.Exit(1)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[INFO], fmt.Sprintf(format, args...)))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[WARNING], fmt.Sprint(args...)))
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[DEBUG], fmt.Sprint(args...)))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[ERROR], fmt.Sprint(args...)))
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *logger) Fatalf(format string, args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", levelFlags[FATAL], fmt.Sprint(args...)))
}
