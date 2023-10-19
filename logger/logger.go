package logger

// A Level is a logging priority. Higher levels are more important.
type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

type ILogger interface {
	Info(args ...any)
	Warn(args ...any)
	Debug(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Debugf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

var (
	levelFlags = []string{"DEBG", "INFO", "WARN", "ERRO", "FATL"}
)

func LevelString(lv Level) string {
	return levelFlags[lv]
}

var (
	log ILogger
)

func SetDefaultLogger(logger ILogger) {
	log = logger
}

func Info(args ...any) {
	log.Info(args...)
}

func Warn(args ...any) {
	log.Warn(args...)
}

func Debug(args ...any) {
	log.Debug(args...)
}

func Error(args ...any) {
	log.Error(args...)
}

func Fatal(args ...any) {
	log.Error(args...)
}

func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...any) {
	log.Warnf(format, args...)
}

func Debugf(format string, args ...any) {
	log.Debugf(format, args...)
}

func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
	log.Fatalf(format, args...)
}
