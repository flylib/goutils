package zaplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func NewZapLogger(options ...Options) *zap.Logger {
	conf := &Config{}
	for i := 0; i < len(options); i++ {
		options[i](conf)
	}

	encoderConfig := GenEncoderConfig(conf)

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //错误等级格式输出样式

	var cores []zapcore.Core

	//时间格式
	if conf.timeLayout == "" {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(conf.timeLayout)
	}

	//文件同步
	if conf.Filename != "" {
		cores = append(cores,
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), FileSyncer(conf), conf.level))
	}

	//控制台同步
	if conf.ENV == DEVELOPMENT || cores == nil {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cores = append(cores,
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), ConsoleSyncer(), conf.level))
	}
	return zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}

//控制台同步
func ConsoleSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
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

func GenEncoderConfig(conf *Config) zapcore.EncoderConfig {
	if conf.ENV == DEVELOPMENT {
		return zap.NewDevelopmentEncoderConfig()
	}
	return zap.NewProductionEncoderConfig()
}
