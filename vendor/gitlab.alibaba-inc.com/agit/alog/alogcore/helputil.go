package alogcore

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
)

var filePath string

func init() {
	filePath = "./logs/main.log"
}

// Get the log content from log file
func getLogContent() (string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}
	contentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(contentBytes), nil
}

func clearLogContent() error {
	err := ioutil.WriteFile(filePath, nil, 0666)
	if err != nil {
		return err
	}
	return nil
}

func checkLogField(assert *assert.Assertions, logContent string, field Field) {
	switch field.zapField.Type {
	case zapcore.StringType:
		assert.Contains(logContent, fmt.Sprintf("\"%s\":\"%s\"", field.zapField.Key, field.zapField.String))
	case zapcore.BoolType:
		assert.Contains(logContent, fmt.Sprintf("\"%s\":%v", field.zapField.Key, field.zapField.Integer == 1))
	default:
		assert.Contains(logContent, fmt.Sprintf("\"%s\":%v", field.zapField.Key, field.zapField.Integer))
	}
}

func checkMsgAndTraceID(assert *assert.Assertions, logContent string, msg string) {
	assert.Contains(logContent, fmt.Sprintf("\"%s\":\"%s\"", JaegerTraceID.GetValue(), "123321"))
	assert.Contains(logContent, fmt.Sprintf("\"%s\":\"%s\"", "msg", msg))
}

func checkLogLevel(assert *assert.Assertions, logContent string, level zapcore.Level) {
	assert.Contains(logContent, fmt.Sprintf("\"%s\":\"%s\"", "level", level.String()))
}

func testCommonLog(t *testing.T, level zapcore.Level) {
	assert := assert.New(t)
	//ctx := context.WithValue(context.Background(), JaegerTraceID, "123321")
	pair := metadata.Pairs(JaegerTraceID.GetValue(), "123321")
	ctx := metadata.NewIncomingContext(context.Background(), pair)
	levelName := level.String()
	type args struct {
		ctx    context.Context
		msg    string
		fields []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: fmt.Sprintf("Test %s", levelName),
			args: args{
				ctx:    ctx,
				msg:    fmt.Sprintf("this is test for %s", levelName),
				fields: []Field{}},
		}, {
			name: fmt.Sprintf("Test %s with field", levelName),
			args: args{
				ctx: ctx,
				msg: fmt.Sprintf("this is test for %s with field", levelName),
				fields: []Field{NewStringField("stringfield", "stringfieldvalue"),
					NewBoolField("boolfield", true)}},
		}, {
			name: fmt.Sprintf("Test %s with int field", levelName),
			args: args{
				ctx: ctx,
				msg: fmt.Sprintf("this is test for %s with int type field", levelName),
				fields: []Field{NewInt64Field("int64", 9988998),
					NewIntField("int", 90909090)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearLogContent()
			switch level {
			case zapcore.DebugLevel:
				Debug(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.InfoLevel:
				Info(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.WarnLevel:
				Warn(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.ErrorLevel:
				Error(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.FatalLevel:
				Fatal(tt.args.ctx, tt.args.msg, tt.args.fields...)
			}
			logContent, err := getLogContent()
			assert.NoError(err)
			checkMsgAndTraceID(assert, logContent, tt.args.msg)
			checkLogLevel(assert, logContent, level)
			for _, field := range tt.args.fields {
				checkLogField(assert, logContent, field)
			}
		})
	}
}

func testLogWithField(t *testing.T, level zapcore.Level) {
	assert := assert.New(t)
	ctx := context.WithValue(context.Background(), JaegerTraceID, "123321")
	levelName := level.String()
	type args struct {
		ctx    context.Context
		msg    string
		fields []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: fmt.Sprintf("Test %s", levelName),
			args: args{
				ctx:    ctx,
				msg:    fmt.Sprintf("this is test for %s", levelName),
				fields: []Field{}},
		}, {
			name: fmt.Sprintf("Test %s with field", levelName),
			args: args{
				ctx: ctx,
				msg: fmt.Sprintf("this is test for %s with field", levelName),
				fields: []Field{NewStringField("stringfield", "stringfieldvalue"),
					NewBoolField("boolfield", true)}},
		}, {
			name: fmt.Sprintf("Test %s with int field", levelName),
			args: args{
				ctx: ctx,
				msg: fmt.Sprintf("this is test for %s with int type field", levelName),
				fields: []Field{NewInt64Field("int64", 9988998),
					NewIntField("int", 90909090)},
			},
		},
	}

	clearLogContent()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch level {
			case zapcore.DebugLevel:
				Debug(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.InfoLevel:
				Info(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.WarnLevel:
				Warn(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.ErrorLevel:
				Error(tt.args.ctx, tt.args.msg, tt.args.fields...)
			case zapcore.FatalLevel:
				Fatal(tt.args.ctx, tt.args.msg, tt.args.fields...)
			}
			logContent, err := getLogContent()
			assert.NoError(err)
			checkMsgAndTraceID(assert, logContent, tt.args.msg)
			checkLogLevel(assert, logContent, level)
			for _, field := range tt.args.fields {
				checkLogField(assert, logContent, field)
			}
		})
	}
}
