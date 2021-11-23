package alogcore

import (
	"context"
	"fmt"
)

// WithFieldLogger type for write fields first
type WithFieldLogger struct {
	fields []Field
}

// WithField write log with fields, fieldValue is a interface
func WithField(fieldName string, fieldValue interface{}) *WithFieldLogger {
	field := newFieldWithInterface(fieldName, fieldValue)
	withFieldLogger := &WithFieldLogger{}
	withFieldLogger.setField(field)
	return withFieldLogger
}

// WithError is write error message and stacktrace
func WithError(err error) *WithFieldLogger {
	if err == nil {
		return &WithFieldLogger{}
	}

	errMsg := NewStringField("ErrorMsg", err.Error())
	errStack := NewStringField("StackTrace", fmt.Sprintf("%+v", err))
	withFieldLogger := &WithFieldLogger{}
	withFieldLogger.setField(errMsg)
	withFieldLogger.setField(errStack)
	return withFieldLogger
}

func (w *WithFieldLogger) setField(field Field) {
	w.fields = append(w.fields, field)
}

func (w *WithFieldLogger) setFields(fields []Field) {
	w.fields = append(w.fields, fields...)
}

// Info write info log with WithFieldLogger
func (w *WithFieldLogger) Info(ctx context.Context, msg string, field ...Field) {
	w.setFields(field)
	log.Info(ctx, msg, w.fields...)
}

// Debug write debug log with WithFieldLogger
func (w *WithFieldLogger) Debug(ctx context.Context, msg string, field ...Field) {
	w.setFields(field)
	log.Debug(ctx, msg, w.fields...)
}

// Warn write warn log with WithFieldLogger
func (w *WithFieldLogger) Warn(ctx context.Context, msg string, field ...Field) {
	w.setFields(field)
	log.Warn(ctx, msg, w.fields...)
}

// Error write error log with WithFieldLogger
func (w *WithFieldLogger) Error(ctx context.Context, msg string, field ...Field) {
	w.setFields(field)
	log.Error(ctx, msg, w.fields...)
}

// Panic write panic log with WithFieldLogger
func (w *WithFieldLogger) Panic(ctx context.Context, msg string, field ...Field) {
	w.setFields(field)
	log.Panic(ctx, msg, w.fields...)
}

// Fatal write fatal log with WithFieldLogger
func (w *WithFieldLogger) Fatal(ctx context.Context, msg string, field ...Field) {
	w.setFields(field)
	log.Fatal(ctx, msg, w.fields...)
}
