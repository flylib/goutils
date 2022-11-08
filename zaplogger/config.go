package zaplogger

import "go.uber.org/zap/zapcore"

type Options func(o *Config)

type Config struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`

	//development or production.The default is DEVELOPMENT
	ENV int8 `json:"env"`

	level zapcore.Level

	//default ISO8601
	timeLayout string
}

const (
	DEVELOPMENT = 0
	PRODUCTION  = 1
)

//同步日志文件
func WithSyncFile(file string) Options {
	return func(o *Config) {
		o.Filename = file
	}
}

//日志文件大小分割界限（M）,默认100Mb
func WithRotatedMaxSize(maxFileSize int) Options {
	return func(o *Config) {
		o.MaxSize = maxFileSize
	}
}

//日志文件保最大留时长分(天),默认不移除
func WithMaxAge(maxFileSize int) Options {
	return func(o *Config) {
		o.MaxSize = maxFileSize
	}
}

//日志文件保最大留时长分(天),默认不移除
func WithTimeLayout(layout string) Options {
	return func(o *Config) {
		o.timeLayout = layout
	}
}

//生产模式，只输出到日志文件
func WithProduction() Options {
	return func(o *Config) {
		o.ENV = PRODUCTION
	}
}
