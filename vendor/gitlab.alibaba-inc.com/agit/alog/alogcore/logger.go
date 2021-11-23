package alogcore

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
)

type logger struct {
	log *zap.Logger
}

var log ALogger

// It will init a default log
func init() {
	InitLogger("/tmp/galileo/unit_test/test.log",
		zap.DebugLevel,
		10,
		1,
		1,
		false,
		false,
		"GalileoDefaultLog")
}

// InitLogger use to init log
func InitLogger(filePath string, level zapcore.Level, maxSize int, maxBackup int, maxAge int, compress bool, stdout bool, serviceName string) {
	log = initLogger(filePath, level, maxSize, maxBackup, maxAge, compress, stdout, serviceName)
}

func initLogger(filePath string, level zapcore.Level, maxSize int, maxBackup int, maxAge int, compress bool, stdout bool, serviceName string) *logger {
	log := &logger{}

	log.log = NewLogger(filePath, /* file Path */
		level,     /* log level */
		maxSize,   /* max size*/
		maxBackup, /* max backup */
		maxAge,    /* max age */
		compress,  /* is compress? */
		stdout,
		serviceName /* service name */)
	return log
}

func (logger *logger) Info(ctx context.Context, msg string, fields ...Field) {
	zapFields := processContext(ctx, fields...)
	logger.log.Info(msg, zapFields...)
}

func (logger *logger) Debug(ctx context.Context, msg string, fields ...Field) {
	zapFields := processContext(ctx, fields...)
	logger.log.Debug(msg, zapFields...)
}

func (logger *logger) Warn(ctx context.Context, msg string, fields ...Field) {
	zapFields := processContext(ctx, fields...)
	logger.log.Warn(msg, zapFields...)
}

func (logger *logger) Error(ctx context.Context, msg string, fields ...Field) {
	zapFields := processContext(ctx, fields...)
	logger.log.Error(msg, zapFields...)
}

func (logger *logger) Panic(ctx context.Context, msg string, fields ...Field) {
	zapFields := processContext(ctx, fields...)
	logger.log.Panic(msg, zapFields...)
}

func (logger *logger) Fatal(ctx context.Context, msg string, fields ...Field) {
	zapFields := processContext(ctx, fields...)
	logger.log.Fatal(msg, zapFields...)
}

// Info will write info type log
// Context will contains other information to log
func Info(ctx context.Context, msg string, fields ...Field) {
	log.Info(ctx, msg, fields...)
}

// Debug will write debug type log
// Context will contains other information to log
func Debug(ctx context.Context, msg string, fields ...Field) {
	log.Debug(ctx, msg, fields...)
}

// Warn will write warn type log
// Context will contains other information to log
func Warn(ctx context.Context, msg string, fields ...Field) {
	log.Warn(ctx, msg, fields...)
}

// Error will write error type log
// Context will contains other information to log
func Error(ctx context.Context, msg string, fields ...Field) {
	log.Error(ctx, msg, fields...)
}

// Panic will write panic type log
// Context will contains other information to log
func Panic(ctx context.Context, msg string, fields ...Field) {
	log.Panic(ctx, msg, fields...)
}

// Fatal will write fatal type log
// Context will contains other information to log
func Fatal(ctx context.Context, msg string, fields ...Field) {
	log.Fatal(ctx, msg, fields...)
}

// GetLogger Get alog logger
func GetLogger() ALogger {
	return log
}

// Add traceId to fields
func processContext(ctx context.Context, fields ...Field) []zap.Field {
	var zapField []zap.Field

	// Start process eagle eye traceID
	if eagleField, ok := getTraceIDFromContext(ctx, EagleEyeTraceID.GetValue()); ok {
		fields = append(fields, eagleField)
	}

	// Start process common traceID where store in ctx string
	if jaegerField, ok := getTraceIDFromContext(ctx, JaegerTraceID.GetValue()); ok {
		fields = append(fields, jaegerField)
	}

	for _, f := range fields {
		zapField = append(zapField, f.zapField)
	}

	return zapField
}

func getTraceIDFromContext(ctx context.Context, key string) (_ Field, ok bool) {
	m, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if v, ok := m[key]; ok {
			if len(v) > 0 {
				eagleTraceField := NewStringField(key, v[0])
				return eagleTraceField, true
			}
		}
	}

	return Field{}, false
}
