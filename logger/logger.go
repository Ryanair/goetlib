package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger(logLevel string, application string, function string, logGroup string, logStream string, vpc string,
	env string) {
	zap.NewProductionConfig()
	rawLog, _ := buildConfig(logLevel).Build()
	rawLog = rawLog.
		WithOptions(zap.AddCallerSkip(1)).
		With(zap.String("application", application)).
		With(zap.String("function_name", function)).
		With(zap.String("@log_group", logGroup)).
		With(zap.String("@log_stream", logStream)).
		With(zap.String("@vpc", vpc)).
		With(zap.String("@env", env))
	Log = rawLog.Sugar()
}

func InitLoggerTest() {
	InitLogger("debug", "test", "function", "log_group", "log_stream", "vpc", "env")
}

func Debug(msg string, vars ...interface{}) {
	Log.Debugf(msg, vars...)
}

func Info(msg string, vars ...interface{}) {
	Log.Infof(msg, vars...)
}

func Warn(msg string, vars ...interface{}) {
	Log.Warnf(msg, vars...)
}

func Error(msg string, vars ...interface{}) {
	Log.Errorf(msg, vars...)
}

func AddToContext(k string, v string) {
	if v != "" {
		Log = Log.With(zap.String(k, v))
	}
}

func AddRefToContext(k string, v *string) {
	if v != nil {
		Log = Log.With(zap.String(k, *v))
	}
}

func buildConfig(logLevel string) zap.Config {
	return zap.Config{
		Level:       readLogLevel(logLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    buildEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func readLogLevel(logLevel string) zap.AtomicLevel {
	logAtomicLevel := zap.AtomicLevel{}
	if err := logAtomicLevel.UnmarshalText([]byte(logLevel)); err != nil {
		fmt.Printf("malformed log level: %+v\n", logLevel)
		logAtomicLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	return logAtomicLevel
}

func buildEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "logger_name",
		MessageKey:     "message",
		StacktraceKey:  "stack_trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
