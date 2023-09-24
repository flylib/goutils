package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Logger struct {
	syncFile    string //同步写入文件
	syncConsole bool   //控制台
	syncEmail   string //同步邮箱
	syncHttp    string //同步http
	*zap.SugaredLogger
}

func WithSyncWriteFile(file string) func(logger *Logger) {
	return func(logger *Logger) {
		logger.syncFile = file
	}
}
func WithSyncConsole(ok bool) func(logger *Logger) {
	return func(logger *Logger) {
		logger.syncConsole = ok
	}
}

func NewZapLogger(options ...func(logger *Logger)) *Logger {
	l := Logger{}
	for _, f := range options {
		f(&l)
	}

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
	}
	encoder := zapcore.EncoderConfig{
		CallerKey:      "caller", // 打印文件名和行数
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
	if l.syncConsole {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel)))
	}
	if l.syncFile != "" {
		syncFile := zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.syncFile, // ⽇志⽂件路径
			MaxSize:    1 << 30,    // 1gb 单位为MB,默认为512MB
			MaxBackups: 7,          // 保留旧文件最大个数
			MaxAge:     3,          // 文件最多保存多少天
			LocalTime:  true,       // 采用本地时间
			Compress:   false,      // 文件是否压缩
		})
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncFile, zap.NewAtomicLevelAt(zap.InfoLevel)))
	}

	zapplog := zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(1))
	return &Logger{SugaredLogger: zapplog.Sugar()}
}

func (l *Logger) Info(msg string, fields ...any) {
	l.Infof(msg, fields)
}

func (l *Logger) Warn(msg string, fields ...any) {
	l.Warnf(msg, fields)
}

func (l *Logger) Debug(msg string, fields ...any) {
	l.Debugf(msg, fields)
}

func (l Logger) Error(msg string, fields ...any) {
	l.Errorf(msg, fields)
}

func (l *Logger) Fatal(msg string, fields ...any) {
	l.Fatalf(msg, fields)
}
