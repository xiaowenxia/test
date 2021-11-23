package alogcore

import (
	"context"
)

type ALogger interface {
	Debug(context.Context, string, ...Field)
	Info(context.Context, string, ...Field)
	Warn(context.Context, string, ...Field)
	Error(context.Context, string, ...Field)
	Panic(context.Context, string, ...Field)
	Fatal(context.Context, string, ...Field)
}
