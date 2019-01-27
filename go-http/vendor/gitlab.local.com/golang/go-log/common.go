// Package golog common variable/const/type/struct/func ...
// Created by chenguolin 2018-12-26
package go_log

// EncoderType define
type EncoderType int

const (
	// JSONEncoder Defined JSON encoder type for logger
	JSONEncoder EncoderType = iota
	// ConsoleEncoder Defined Console encoder type for logger
	ConsoleEncoder
)

// LogLevel define
type LogLevel int

const (
	// DebugLevel debug
	DebugLevel LogLevel = iota
	// InfoLevel info
	InfoLevel
	// WarnLevel warn
	WarnLevel
	// ErrorLevel error
	ErrorLevel
	// PanicLevel panic
	PanicLevel
	// FatalLevel fatal
	FatalLevel
)

// const string
const (
	// NameZapJSONLogger ZapJSONLogger
	NameZapJSONLogger = "ZapJSONLogger"
	// NameConsoleLogger ZapConsoleLogger
	NameConsoleLogger = "ZapConsoleLogger"

	// NameKey Name
	NameKey = "Name"
	// StacktraceKey Stack
	StacktraceKey = "Stack"
	// MessageKey Message
	MessageKey = "Message"
	// LevelKey level
	LevelKey = "level"
	// TimeKey  Time
	TimeKey = "TimeStamp"
	// CallerKey  Caller
	CallerKey = "Caller"

	// DefaultTimePattern time pattern
	DefaultTimePattern = "2006-01-02-15-04-05"
)
