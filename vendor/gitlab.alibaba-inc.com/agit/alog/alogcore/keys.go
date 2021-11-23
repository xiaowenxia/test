package alogcore

// TraceIDKey trace ID will write to log
type jaegerTraceIDKey string

// JaegerTraceID the traceID
var JaegerTraceID jaegerTraceIDKey = "x-jaeger-trace-id-galileo"

// GetValue Get TraceIDKey string value
func (t jaegerTraceIDKey) GetValue() string {
	return string(t)
}

// Eagle eye trace ID, it come from force
type eagleEyeTraceIDKey string

// EagleEyeTraceID the eagle eye traceID
var EagleEyeTraceID eagleEyeTraceIDKey = "x-eagle-eye-trace-id"

// GetValue get EagleEyeTraceID string value
func (e eagleEyeTraceIDKey) GetValue() string {
	return string(e)
}
