package zaplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func NewZapLogger(options ...Options) *zap.SugaredLogger {
	conf := &Config{}
	for i := 0; i < len(options); i++ {
		options[i](conf)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var cores []zapcore.Core

	//文件同步
	if conf.Filename != "" {
		cores = append(cores,
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), FileSyncer(conf), conf.Level))
	}
	//开发模式控制台输出
	if conf.ENV == DEVELOPMENT || cores == nil {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cores = append(cores,
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), conf.Level))
	}
	Logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	return Logger.Sugar()
}

//文件同步
func FileSyncer(conf *Config) zapcore.WriteSyncer {
	fileLogger := &lumberjack.Logger{
		Filename:  conf.Filename,
		MaxSize:   conf.MaxSize, // 1gb 单位为MB,默认为512MB
		MaxAge:    conf.MaxAge,
		Compress:  false,
		LocalTime: true, //是否按当地（本机）时间重命名备份文件
	}
	return zapcore.AddSync(fileLogger)
}
