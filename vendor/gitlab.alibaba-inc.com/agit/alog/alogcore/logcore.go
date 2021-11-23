package alogcore

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger create logger type
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackup int, maxAge int, compress bool, stdout bool,
	serviceName string) *zap.Logger {
	core := newZapCore(filePath, level, maxSize, maxBackup, maxAge, compress, stdout)
	return zap.New(core, zap.AddCaller(),
		zap.Development(),
		zap.Fields(zap.String("serviceName", serviceName)),
		zap.AddCallerSkip(2))
}

func newZapCore(filePath string, level zapcore.Level, maxSize int, maxBackup int, maxAge int, compress bool, stdout bool) zapcore.Core {
	logRoll := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   compress,
		LocalTime:  true,
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	zapEncoder := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "lineNume",
		StacktraceKey:  "Alitrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	writeSyncer := make([]zapcore.WriteSyncer, 0, 2)
	writeSyncer = append(writeSyncer, zapcore.AddSync(&logRoll))
	if stdout {
		writeSyncer = append(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(zapEncoder),
		zapcore.NewMultiWriteSyncer(writeSyncer...),
		atomicLevel,
	)
}
