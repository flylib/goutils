package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type option struct {
	syncInfoFile, syncErrorFile string //同步写入文件
	syncConsole                 bool   //控制台
	syncEmail                   string //同步邮箱
	syncHttp                    string //同步http
	timeformat                  string //时间格式
}

func NewZapLogger(options ...Option) *zap.SugaredLogger {
	l := option{}
	for _, f := range options {
		f(&l)
	}

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		if l.timeformat == "" {
			l.timeformat = "2006-01-02 15:04:05"
		}
		enc.AppendString(t.Format(l.timeformat))
	}
	encoder := zapcore.EncoderConfig{
		CallerKey:      "caller", // 打印文件名和行数 json格式时生效
		LevelKey:       "lv",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,                // 自定义时间格式
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder,       // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var cores []zapcore.Core
	if l.syncConsole || (l.syncInfoFile == "" && l.syncErrorFile == "") {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel)))
	}
	if l.syncInfoFile != "" {
		level := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		})
		syncFile := zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.syncInfoFile, // ⽇志⽂件路径
			MaxSize:    1 << 30,        // 1gb 单位为MB,默认为512MB
			MaxBackups: 7,              // 保留旧文件最大个数
			MaxAge:     3,              // 文件最多保存多少天
			LocalTime:  true,           // 采用本地时间
			Compress:   false,          // 文件是否压缩
		})
		encoder.EncodeLevel = zapcore.CapitalLevelEncoder
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncFile, level))
	}

	if l.syncErrorFile != "" {
		level := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
			return lev >= zap.ErrorLevel
		})
		syncFile := zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.syncErrorFile, // ⽇志⽂件路径
			MaxSize:    1 << 30,         // 1gb 单位为MB,默认为512MB
			MaxBackups: 7,               // 保留旧文件最大个数
			MaxAge:     3,               // 文件最多保存多少天
			LocalTime:  true,            // 采用本地时间
			Compress:   false,           // 文件是否压缩
		})
		encoder.EncodeLevel = zapcore.CapitalLevelEncoder
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncFile, level))
	}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller()).Sugar()
}
