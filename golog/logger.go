/*
Package golog zap日志封装
Created by chenguolin 2018-12-26

github: https://github.com/uber-go/zap
zap doc: https://godoc.org/go.uber.org/zap
*/
package golog

import (
	"errors"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger use zap logger
type Logger struct {
	Log *zap.Logger
}

// NewLogger new logger with the given Options
// default use InfoLevel, JSONEncoder, os.Stdout
func NewLogger(opts ...Option) (*Logger, error) {
	// default config out to os.Stdout
	cfg := &Config{
		Level:       InfoLevel,
		Encoder:     JSONEncoder,
		WithCaller:  false,
		Out:         os.Stdout,
		TimePattern: DefaultTimePattern,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// NewZapLogger
	// 默认使用ZapLogger
	return newZapLogger(cfg)
}

// newZapLogger use config new zap logger
func newZapLogger(c *Config) (*Logger, error) {
	// 1. set zap level
	lv, err := setLogLevel(c.Level)
	if err != nil {
		return nil, err
	}

	// 2. set zap output
	var output zapcore.WriteSyncer
	if c.Out == nil {
		return nil, errors.New("Error Config.Out is nil")
	}
	// set output
	output = zapcore.AddSync(c.Out)
	if !c.WithNoLock {
		output = zapcore.Lock(output)
	}

	// 3. set zapEncoder
	zapEncoder, name, err := setEncoder(c.TimePattern, c.Encoder)
	if err != nil {
		return nil, err
	}

	// 4. new zap logger
	var zapOpts []zap.Option
	log := zap.New(zapcore.NewCore(zapEncoder, output, lv), zapOpts...)

	// 5. return logger
	logger := &Logger{
		Log: log.Named(name),
	}
	return logger, nil
}

// setLogLevel set log level
func setLogLevel(level LogLevel) (zap.AtomicLevel, error) {
	var lv = zap.NewAtomicLevel()

	// set level
	switch level {
	case DebugLevel:
		lv.SetLevel(zap.DebugLevel)
	case InfoLevel:
		lv.SetLevel(zap.InfoLevel)
	case WarnLevel:
		lv.SetLevel(zap.WarnLevel)
	case ErrorLevel:
		lv.SetLevel(zap.ErrorLevel)
	case PanicLevel:
		lv.SetLevel(zap.PanicLevel)
	case FatalLevel:
		lv.SetLevel(zap.FatalLevel)
	default:
		return lv, errors.New("Error setLogLevel log level invalid")
	}

	return lv, nil
}

// setEncoder set encoder type
func setEncoder(timePattern string, encoderType EncoderType) (zapcore.Encoder, string, error) {
	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format(timePattern))
	}
	encoderCfg := zapcore.EncoderConfig{
		NameKey:        NameKey,
		StacktraceKey:  StacktraceKey,
		MessageKey:     MessageKey,
		LevelKey:       LevelKey,
		TimeKey:        TimeKey,
		CallerKey:      CallerKey,
		EncodeTime:     timeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// set zap encoder
	var name string
	var zapEncoder zapcore.Encoder

	switch encoderType {
	case JSONEncoder:
		zapEncoder = zapcore.NewJSONEncoder(encoderCfg)
		name = NameZapJSONLogger
	case ConsoleEncoder:
		zapEncoder = zapcore.NewConsoleEncoder(encoderCfg)
		name = NameConsoleLogger
	default:
		return zapEncoder, name, errors.New("Error setEncoder encoder type invalid")
	}

	return zapEncoder, name, nil
}

// Debug log
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Log.Debug(msg, fields...)
}

// Info log
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Log.Info(msg, fields...)
}

// Warn log
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Log.Warn(msg, fields...)
}

// Error log
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Log.Error(msg, fields...)
}

// Panic log
func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.Log.Panic(msg, fields...)
}

// Fatal log
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Log.Fatal(msg, fields...)
}
