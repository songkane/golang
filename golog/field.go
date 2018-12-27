/*
Package golog zapcore.Field函数封装
Created by chenguolin 2018-12-26
*/
package golog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Bool constructs a field that carries a bool.
func Bool(key string, val bool) zapcore.Field {
	return zap.Bool(key, val)
}

// Bools constructs a field that carries a slice of bools.
func Bools(key string, vars []bool) zapcore.Field {
	return zap.Bools(key, vars)
}

// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
// To log opaque binary blobs (which aren't necessarily valid UTF-8), use Binary.
func ByteString(key string, val []byte) zapcore.Field {
	return zap.ByteString(key, val)
}

// ByteStrings constructs a field that carries a slice of []byte, each of which must be UTF-8 encoded text.
func ByteStrings(key string, vals [][]byte) zapcore.Field {
	return zap.ByteStrings(key, vals)
}

// Duration constructs a field with the given key and value.
// The encoder controls how the duration is serialized.
func Duration(key string, val time.Duration) zapcore.Field {
	return zap.Duration(key, val)
}

// Durations constructs a field that carries a slice of time.Durations.
func Durations(key string, vals []time.Duration) zapcore.Field {
	return zap.Durations(key, vals)
}

// Err is shorthand for the common idiom NamedError("error", err).
func Err(err error) zapcore.Field {
	return zap.Error(err)
}

// Errs constructs a field that carries a slice of errors.
func Errs(key string, errs []error) zapcore.Field {
	return zap.Errors(key, errs)
}

// Float32 constructs a field that carries a float32.
// The way the floating-point value is represented is encoder-dependent, so marshaling is necessarily lazy.
func Float32(key string, val float32) zapcore.Field {
	return zap.Float32(key, val)
}

// Float32s constructs a field that carries a slice of floats.
func Float32s(key string, vals []float32) zapcore.Field {
	return zap.Float32s(key, vals)
}

// Float64 constructs a field that carries a float64.
// The way the floating-point value is represented is encoder-dependent, so marshaling is necessarily lazy.
func Float64(key string, val float64) zapcore.Field {
	return zap.Float64(key, val)
}

// Float64s constructs a field that carries a slice of floats.
func Float64s(key string, vals []float64) zapcore.Field {
	return zap.Float64s(key, vals)
}

// Int8 constructs a Field with the given key and value.
func Int8(key string, val int8) zapcore.Field {
	return zap.Int8(key, val)
}

// Int8s constructs a Field with the given key and value.
func Int8s(key string, vals []int8) zapcore.Field {
	return zap.Int8s(key, vals)
}

// Int16 constructs a field with the given key and value.
func Int16(key string, val int16) zapcore.Field {
	return zap.Int16(key, val)
}

// Int16s constructs a field that carries a slice of integers.
func Int16s(key string, vals []int16) zapcore.Field {
	return zap.Int16s(key, vals)
}

// Int32 constructs a field with the given key and value.
func Int32(key string, val int32) zapcore.Field {
	return zap.Int32(key, val)
}

// Int32s constructs a field that carries a slice of integers.
func Int32s(key string, vals []int32) zapcore.Field {
	return zap.Int32s(key, vals)
}

// Int64 constructs a field with the given key and value.
func Int64(key string, val int64) zapcore.Field {
	return zap.Int64(key, val)
}

// Int64s constructs a field that carries a slice of integers.
func Int64s(key string, vals []int64) zapcore.Field {
	return zap.Int64s(key, vals)
}

// Int constructs a field with the given key and value.
func Int(key string, val int) zapcore.Field {
	return zap.Int(key, val)
}

// Ints constructs a field that carries a slice of integers.
func Ints(key string, vals []int) zapcore.Field {
	return zap.Ints(key, vals)
}

// Uint8 constructs a field with the given key and value.
func Uint8(key string, val uint8) zapcore.Field {
	return zap.Uint8(key, val)
}

// Uint8s constructs a field that carries a slice of integers.
func Uint8s(key string, vals []uint8) zapcore.Field {
	return zap.Uint8s(key, vals)
}

// Uint16 constructs a field with the given key and value.
func Uint16(key string, val uint16) zapcore.Field {
	return zap.Uint16(key, val)
}

// Uint16s constructs a field that carries a slice of integers.
func Uint16s(key string, vals []uint16) zapcore.Field {
	return zap.Uint16s(key, vals)
}

// Uint32 constructs a field with the given key and value.
func Uint32(key string, val uint32) zapcore.Field {
	return zap.Uint32(key, val)
}

// Uint32s constructs a field that carries a slice of integers.
func Uint32s(key string, vals []uint32) zapcore.Field {
	return zap.Uint32s(key, vals)
}

// Uint64 constructs a field with the given key and value.
func Uint64(key string, val uint64) zapcore.Field {
	return zap.Uint64(key, val)
}

// Uint64s constructs a field that carries a slice of integers.
func Uint64s(key string, vals []uint64) zapcore.Field {
	return zap.Uint64s(key, vals)
}

// Uint constructs a Field with the given key and value.
func Uint(key string, val uint) zapcore.Field {
	return zap.Uint(key, val)
}

// Uints constructs a Field with the given key and value.
func Uints(key string, vals []uint) zapcore.Field {
	return zap.Uints(key, vals)
}

// Object constructs a field with the given key and ObjectMarshaler.
// It provides a flexible, but still type-safe and efficient, way to add map- or struct-like user-defined types to the logging context.
// The struct's MarshalLogObject method is called lazily.
func Object(key string, val interface{}) zapcore.Field {
	return zap.Any(key, val)
}

// Stack constructs a field that stores a stacktrace of the current goroutine under provided key.
// Keep in mind that taking a stacktrace is eager and expensive (relatively speaking);
// this function both makes an allocation and takes about two microseconds.
func Stack(key string) zapcore.Field {
	return zap.Stack(key)
}

// String constructs a field with the given key and value.
func String(key string, val string) zapcore.Field {
	return zap.String(key, val)
}

// Strings constructs a field that carries a slice of strings.
func Strings(key string, val []string) zapcore.Field {
	return zap.Strings(key, val)
}

// Time constructs a Field with the given key and value. The encoder controls how the time is serialized.
func Time(key string, val time.Time) zapcore.Field {
	return zap.Time(key, val)
}

// Times constructs a field that carries a slice of time.Times.
func Times(key string, vals []time.Time) zapcore.Field {
	return zap.Times(key, vals)
}
