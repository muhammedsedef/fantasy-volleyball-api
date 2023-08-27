package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"path"
	"runtime"
)

type LogrusLogger struct {
	name             string
	underlyingLogger *logrus.Logger
	defaultFields    map[string]interface{}
}

func (l LogrusLogger) GetName() string {
	return l.name
}

func (l LogrusLogger) SetOut(o io.Writer) {
	l.underlyingLogger.Out = o
}

func (l LogrusLogger) SetFormatter(formatter logrus.Formatter) {
	l.underlyingLogger.Formatter = formatter
}

func (l LogrusLogger) GetFormatter() logrus.Formatter {
	return l.underlyingLogger.Formatter
}

func (l LogrusLogger) Debug(format string, args ...interface{}) {
	if !isDebugEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Debug(format)
}

func (l LogrusLogger) DebugWithContext(ctx context.Context, format string, args ...interface{}) {
	if !isDebugEnabled(l.name) {
		return
	}

	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Debug(format)
}

func (l LogrusLogger) DebugWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isDebugEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Debug(format)
}

func (l LogrusLogger) DebugWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if !isDebugEnabled(l.name) {
		return
	}

	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Debug(format)
}

func (l LogrusLogger) Info(format string, args ...interface{}) {
	if !isInfoEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Info(format)
}

func (l LogrusLogger) InfoWithContext(ctx context.Context, format string, args ...interface{}) {
	if !isInfoEnabled(l.name) {
		return
	}

	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Info(format)
}

func (l LogrusLogger) InfoWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isInfoEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Info(format)
}

func (l LogrusLogger) InfoWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if !isInfoEnabled(l.name) {
		return
	}

	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Info(format)
}

func (l LogrusLogger) Warn(format string, args ...interface{}) {
	if !isWarnEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Warn(format)
}

func (l LogrusLogger) WarnWithContext(ctx context.Context, format string, args ...interface{}) {
	if !isWarnEnabled(l.name) {
		return
	}

	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Warn(format)
}

func (l LogrusLogger) WarnWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isWarnEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}
	mapFieldsToEntry(entry, fields)
	entry.Warn(format)
}

func (l LogrusLogger) WarnWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if !isWarnEnabled(l.name) {
		return
	}

	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}
	mapFieldsToEntry(entry, fields)
	entry.Warn(format)
}

func (l LogrusLogger) Error(format string, args ...interface{}) {
	if !isErrorEnabled(l.name) {
		return
	}
	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Error(format)
}

func (l LogrusLogger) ErrorWithContext(ctx context.Context, format string, args ...interface{}) {
	if !isErrorEnabled(l.name) {
		return
	}
	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Error(format)
}

func (l LogrusLogger) ErrorWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isErrorEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Error(format)
}

func (l LogrusLogger) ErrorWithFieldsAndContext(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	if !isErrorEnabled(l.name) {
		return
	}

	entry := l.newLogEntryWithContext(ctx)

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Error(format)
}

func (l LogrusLogger) IsDebugEnabled() bool {
	return isDebugEnabled(l.name)
}

func (l LogrusLogger) IsInfoEnabled() bool {
	return isInfoEnabled(l.name)
}

func (l LogrusLogger) IsWarnEnabled() bool {
	return isWarnEnabled(l.name)
}

func (l LogrusLogger) IsErrorEnabled() bool {
	return isErrorEnabled(l.name)
}

func (l LogrusLogger) newLogEntry() *logrus.Entry {
	defaultFields := logrus.Fields(l.defaultFields)

	entry := l.underlyingLogger.WithFields(defaultFields)

	entry.Data["logger"] = l.name
	entry.Data["callStack"] = l.getCallStack()

	return entry
}

func (l LogrusLogger) newLogEntryWithContext(ctx context.Context) *logrus.Entry {
	defaultFields := logrus.Fields(l.defaultFields)

	entry := l.underlyingLogger.WithFields(defaultFields)

	entry.Data["logger"] = l.name
	entry.Data["callStack"] = l.getCallStack()

	if ctx != nil {
		correlationId := ctx.Value("X-CorrelationId")
		if correlationId != nil {
			entry.Data["X-CorrelationId"] = correlationId
		}
	}

	return entry
}

func (l LogrusLogger) getCallStack() map[string]interface{} {
	pc, file, line, _ := runtime.Caller(3)
	_, fileName := path.Split(file)
	fullPath := runtime.FuncForPC(pc).Name()

	return map[string]interface{}{
		"fullPath": fullPath,
		"fileName": fileName,
		"line":     line,
	}
}