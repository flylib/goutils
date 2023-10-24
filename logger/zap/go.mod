module github.com/flylib/goutils/logger/zaplog

go 1.18

require (
	github.com/flylib/goutils/logger v0.0.0-20231023014531-4f50a5871c60
	go.uber.org/zap v1.26.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require go.uber.org/multierr v1.11.0 // indirect

replace github.com/flylib/goutils/logger v0.0.0-20231023014531-4f50a5871c60 => ../../logger
