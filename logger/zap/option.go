package zaplog

type Option func(o *option)

// options
type option struct {
	syncFile    string
	syncConsole bool
	//syncEmail   string //同步邮箱
	//syncHttp    string //同步http
	timeFormat  string
	maxFileSize int //日志文件最大多大
	maxAge      int //文件最多保存多少天
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
