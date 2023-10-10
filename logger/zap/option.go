package zaplog

type Option func(o *option)

// options
func SyncInfoFile(file string) Option {
	return func(o *option) {
		o.syncInfoFile = file
	}
}
func SyncErrorFile(file string) Option {
	return func(o *option) {
		o.syncErrorFile = file
	}
}
func SyncConsole(ok bool) Option {
	return func(o *option) {
		o.syncConsole = ok
	}
}
func TimeFormat(format string) Option {
	return func(o *option) {
		o.timeformat = format
	}
}
