package log

import (
	"context"
)

//go:generate mockery --name=Logger
type Logger interface {
	GetName() string

	Debug(msg string, args ...interface{})
	DebugWithContext(ctx context.Context, msg string, args ...interface{})
	DebugWithFields(fields map[string]interface{}, msg string, args ...interface{})
	DebugWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, msg string, args ...interface{})
	Info(msg string, args ...interface{})
	InfoWithContext(ctx context.Context, msg string, args ...interface{})
	InfoWithFields(fields map[string]interface{}, msg string, args ...interface{})
	InfoWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	WarnWithContext(ctx context.Context, msg string, args ...interface{})
	WarnWithFields(fields map[string]interface{}, msg string, args ...interface{})
	WarnWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, msg string, args ...interface{})
	Error(msg string, args ...interface{})
	ErrorWithContext(ctx context.Context, msg string, args ...interface{})
	ErrorWithFields(fields map[string]interface{}, msg string, args ...interface{})
	ErrorWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, msg string, args ...interface{})

	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
}
