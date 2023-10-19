package log

import (
	"fmt"
	"github.com/flylib/goutils/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	callDepth int
}

func NewLogger(options ...Option) *Logger {
	newLogger := new(Logger)
	newLogger.callDepth = 2

	opt := option{}
	for _, f := range options {
		f(&opt)
	}

	//print to console
	if opt.syncConsole || opt.syncFile == "" {
		newLogger.Logger = log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)
	} else {
		newLogger.Logger = log.New(&lumberjack.Logger{
			Filename:  opt.syncFile,
			MaxSize:   opt.maxFileSize,
			MaxAge:    opt.maxAge,
			LocalTime: true,
			Compress:  false,
		}, "", log.Lshortfile|log.LstdFlags)
	}
	return newLogger
}

func (l *Logger) Debug(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.DebugLevel), fmt.Sprint(args...)))
}

func (l *Logger) Info(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.InfoLevel), fmt.Sprint(args...)))
}

func (l *Logger) Warn(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.WarnLevel), fmt.Sprint(args...)))
}

func (l *Logger) Error(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.ErrorLevel), fmt.Sprint(args...)))
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *Logger) Fatal(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.FatalLevel), fmt.Sprint(args...)))
	os.Exit(1)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.DebugLevel), fmt.Sprintf(format, args...)))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.InfoLevel), fmt.Sprintf(format, args...)))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.WarnLevel), fmt.Sprintf(format, args...)))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.ErrorLevel), fmt.Sprintf(format, args...)))
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", logger.LevelString(logger.FatalLevel), fmt.Sprintf(format, args...)))
}
