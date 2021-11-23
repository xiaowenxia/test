package alogcore

import (
	"fmt"

	"go.uber.org/zap"
)

// Field will be add to log
type Field struct {
	zapField zap.Field
}

// NewStringField is to create a string type field
func NewStringField(key, value string) Field {
	f := Field{}
	f.zapField = zap.String(key, value)
	return f
}

// NewBoolField is to create a bool type field
func NewBoolField(key string, value bool) Field {
	f := Field{}
	f.zapField = zap.Bool(key, value)
	return f
}

// NewFloat64Field is to create a float64 type field
func NewFloat64Field(key string, value float64) Field {
	f := Field{}
	f.zapField = zap.Float64(key, value)
	return f
}

// NewFloat32Field is to create float32 type field
func NewFloat32Field(key string, value float32) Field {
	f := Field{}
	f.zapField = zap.Float32(key, value)
	return f
}

// NewIntField is to create int type field
func NewIntField(key string, value int) Field {
	f := Field{}
	f.zapField = zap.Int(key, value)
	return f
}

// NewInt64Field is to create int64 type field
func NewInt64Field(key string, value int64) Field {
	f := Field{}
	f.zapField = zap.Int64(key, value)
	return f
}

// NewField is to create interface type field, efficiency will be lower
func NewField(key string, value interface{}) Field {
	return newFieldWithInterface(key, value)
}

// Use interface to create Field, efficiency will be lower
func newFieldWithInterface(key string, value interface{}) Field {
	field := Field{}
	switch value.(type) {
	case bool:
		field.zapField = zap.Bool(key, value.(bool))
	case float64:
		field.zapField = zap.Float64(key, value.(float64))
	case float32:
		field.zapField = zap.Float32(key, value.(float32))
	case int:
		field.zapField = zap.Int(key, value.(int))
	case int32:
		field.zapField = zap.Int32(key, value.(int32))
	case int16:
		field.zapField = zap.Int16(key, value.(int16))
	case int8:
		field.zapField = zap.Int8(key, value.(int8))
	case string:
		field.zapField = zap.String(key, value.(string))
	case uint64:
		field.zapField = zap.Uint64(key, value.(uint64))
	case uint32:
		field.zapField = zap.Uint32(key, value.(uint32))
	case uint16:
		field.zapField = zap.Uint16(key, value.(uint16))
	case uint8:
		field.zapField = zap.Uint8(key, value.(uint8))
	default:
		field.zapField = zap.String(key, fmt.Sprintf("%v", value))
	}
	return field
}
