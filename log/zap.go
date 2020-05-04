package log

import (
	"context"

	pawctx "github.com/PAWSOME-INDONESIA/paw_kit/ctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	labelRequestID = "requestID"
	labelData      = "data"
)

type ZapLogger struct {
	logger *zap.Logger
}

func zapConfig(debugMode bool) zap.Config {
	cfg := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.DebugLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	if !debugMode {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cfg.Development = false
		// should we disable stacktrace in a non-debug mode?
		cfg.DisableStacktrace = true
	}

	return cfg
}

// New initiate a new zap logger
func New(cmdName string, debugMode bool) (*ZapLogger, error) {
	withFields := zap.Fields(zap.String("command", cmdName))

	logger, err := zapConfig(debugMode).Build(withFields, zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &ZapLogger{logger: logger}, nil
}

// Stop call this function before exiting program
func (zl *ZapLogger) Stop() error {
	return zl.logger.Sync()
}

// Debug write log with DEBUG level
func (zl *ZapLogger) Debug(ctx context.Context, msg string, data interface{}) {
	zl.logger.Debug(msg, fields(ctx, data, nil)...)
}

// Info write log with INFO level
func (zl *ZapLogger) Info(ctx context.Context, msg string, data interface{}) {
	zl.logger.Info(msg, fields(ctx, data, nil)...)
}

// Warn write log with WARN level
func (zl *ZapLogger) Warn(ctx context.Context, msg string, data interface{}, err error) {
	zl.logger.Warn(msg, fields(ctx, data, err)...)
}

// Error write log with ERROR level
func (zl *ZapLogger) Error(ctx context.Context, msg string, data interface{}, err error) {
	zl.logger.Error(msg, fields(ctx, data, err)...)
}

// Panic write log with PANIC level
func (zl *ZapLogger) Panic(ctx context.Context, msg string, data interface{}, err error) {
	zl.logger.Panic(msg, fields(ctx, data, err)...)
}

// Fatal write log with FATAL level
func (zl *ZapLogger) Fatal(ctx context.Context, msg string, data interface{}, err error) {
	zl.logger.Fatal(msg, fields(ctx, data, err)...)
}

func fields(ctx context.Context, data interface{}, err error) []zap.Field {
	// hard code capacity to 3 (requestID, data, and error)
	fields := make([]zap.Field, 0, 3)
	fields = append(fields, zap.String(labelRequestID, pawctx.RequestID(ctx)))
	fields = append(fields, zap.Any(labelData, data))

	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	return fields
}
