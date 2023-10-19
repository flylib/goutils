package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func NewZapLogger(options ...Option) *zap.Logger {
	opt := option{}
	for _, f := range options {
		f(&opt)
	}

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		if opt.timeFormat == "" {
			opt.timeFormat = time.DateTime
		}
		enc.AppendString(t.Format(opt.timeFormat))
	}
	encoder := zapcore.EncoderConfig{
		CallerKey:      "caller", // 打印文件名和行数 json格式时生效
		LevelKey:       "lv",
		MessageKey:     "msg",
		TimeKey:        "time",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,                // 自定义时间格式
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder,       // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var cores []zapcore.Core
	if opt.syncConsole || opt.syncFile == "" {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel)))
	}
	if opt.syncFile != "" {
		level := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		})
		syncFile := zapcore.AddSync(&lumberjack.Logger{
			Filename:  opt.syncFile,
			MaxSize:   opt.maxFileSize,
			MaxAge:    opt.maxAge,
			LocalTime: true,
			Compress:  false,
		})
		encoder.EncodeLevel = zapcore.CapitalLevelEncoder
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncFile, level))
	}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}
