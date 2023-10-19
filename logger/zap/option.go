package zaplog

type Option func(o *option)

// A Level is a logging priority. Higher levels are more important.
type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
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

// options
type option struct {
	syncFile    string
	syncConsole bool
	//syncEmail   string //同步邮箱
	//syncHttp    string //同步http
	timeFormat    string
	maxFileSize   int //日志文件最大多大
	maxAge        int //文件最多保存多少天
	outJsonStyle  bool
	minPrintLevel Level
}

// 同步写入文件
func SyncFile(file string) Option {
	return func(o *option) {
		o.syncFile = file
	}
}

// 是否同步控制台
func SyncConsole() Option {
	return func(o *option) {
		o.syncConsole = true
	}
}

// 时间格式
func TimeFormat(format string) Option {
	return func(o *option) {
		o.timeFormat = format
	}
}

// 单个日志文件大小（单位:MB）
func MaxFileSize(size int) Option {
	return func(o *option) {
		o.maxFileSize = size
	}
}

// 文件最多保留多长时间(单位:Day)
func MaxSaveDuration(day int) Option {
	return func(o *option) {
		o.maxAge = day
	}
}

// 输出jason格式
func JsonFormat() Option {
	return func(o *option) {
		o.outJsonStyle = true
	}
}

// 最低打印日志级别
func MinPrintLevel(lv Level) Option {
	return func(o *option) {
		o.minPrintLevel = lv
	}
}
