package logger

import (
	"github.com/Gocyber-world/navigator-demo/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getBasicEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		// FunctionKey:    "function",
	}
	return config
}

//TODO: 日志级别需要从配置文件读取或者依据环境进行设置

func getBasicLoggerConfig() zap.Config {
	return zap.Config{
		// 输出的最低日志级别
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "json",
		EncoderConfig:     getBasicEncoderConfig(),
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     map[string]interface{}{},
	}
}

func InitLogger() error {
	var logger *zap.Logger
	var err error
	if global.STAGE == "prod" || global.STAGE == "beta" {
		config := getBasicLoggerConfig()
		if logger, err = config.Build(zap.AddCallerSkip(1)); err != nil {
			return err
		}
		zap.ReplaceGlobals(logger)
		zap.S().Debug("zap logger start")
		return nil

	}

	// 其他环境采用预设
	logger, err = zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	zap.S().Debug("zap logger start")
	return nil

}

// 日志的封装, 以及分等级打印不同的字段
// 目前长期保留的只有 全局的logger, WithOptions 会在这个全局 logger 的基础上克隆一个 logger

func Info(msg string, fields ...zapcore.Field) {
	zap.L().Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

func Error(msg string, fields ...zapcore.Field) {
	zap.L().Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args...)
}
