package logger

import (
	"echo-project/config"
	"echo-project/constant"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// Logger - a custom logger that can be used throughout the application

// Declare the logger instance
var loggerInstance *zap.Logger
var once sync.Once // Use sync.Once to make sure the logger is instantiated only once

// InitLogger - initializes the logger, but only once
func InitLogger() {
	once.Do(func() {
		var err error
		logLevel := config.Config().LogLevel
		cfg := zap.NewProductionConfig()
		if logLevel == "debug" {
			cfg = zap.NewDevelopmentConfig()
		}
		cfg.Encoding = "console"
		cfg.EncoderConfig.ConsoleSeparator = constant.DefaultSeparator
		cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		cfg.EncoderConfig.FunctionKey = "function"
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(constant.DefaultUTCLogFormat))
		})

		logger, err := cfg.Build()
		loggerInstance = logger

		if err != nil {
			panic("failed to initialize zap logger: " + err.Error())
		}
	})
}

// GetLogger return the logger instance
func GetLogger() *zap.Logger {

	if loggerInstance == nil {
		InitLogger()
	}
	return loggerInstance
}

// Debug - logs a message at Debug level with the request id
func Debug(rid string, msg string, fields ...interface{}) {
	GetLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(true)).Debug(fmt.Sprintf("[%s] %s", rid, msg), zap.Reflect("data", fields))
}

// Info - logs a message at Info level with the request id
func Info(rid string, msg string, fields ...interface{}) {
	GetLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(true)).Info(fmt.Sprintf("[%s] %s", rid, msg), zap.Reflect("data", fields))
}

// Error - logs a message at Error level with the request id
func Error(rid string, msg string, err ...interface{}) {
	GetLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(true)).Error(fmt.Sprintf("[%s] %s", rid, msg), zap.Reflect("error", err))
}

// Error - logs a message at Error level with the request id
func Warn(rid string, msg string, err ...interface{}) {
	GetLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(true)).Warn(fmt.Sprintf("[%s] %s", rid, msg), zap.Reflect("error", err))
}

// Close flushes any buffered log entries. Processes should take care to call Sync before exiting.
func CloseLogger() {
	if loggerInstance != nil {
		if err := loggerInstance.Sync(); err != nil {
			loggerInstance.Error("Failed to flush log buffer", zap.Error(err))
		}
	}
}
