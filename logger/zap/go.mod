module github.com/flylib/goutils/logger/zap

go 1.18

require (
	go.uber.org/zap v1.26.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/flylib/goutils/logger v0.0.0-20231019065213-26ef683234a6 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)


replace (
	github.com/flylib/goutils/logger v0.0.0-20231019065213-26ef683234a6 => ../../logger
)