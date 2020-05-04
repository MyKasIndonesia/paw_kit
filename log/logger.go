package log

import "context"

// Logger a logger
type Logger interface {
	Stop() error
	Debug(ctx context.Context, msg string, data interface{})
	Info(ctx context.Context, msg string, data interface{})
	Warn(ctx context.Context, msg string, data interface{}, err error)
	Error(ctx context.Context, msg string, data interface{}, err error)
	Panic(ctx context.Context, msg string, data interface{}, err error)
	Fatal(ctx context.Context, msg string, data interface{}, err error)
}
