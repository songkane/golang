/*
Package golog config
Created by chenguolin 2018-12-26
*/
package golog

import (
	"io"
	"os"
)

// Config for logger
type Config struct {
	Level       LogLevel    //logger level
	Encoder     EncoderType //json or console encoder
	WithCaller  bool        //print the fileName & line number within the log
	Out         io.Writer   //log out
	WithNoLock  bool        //wether add the lock for the zap writer, default is false
	TimePattern string      //Define the time pattern for use
}

// Option set the option to Config
type Option func(c *Config)

// WithOutput set config Out
func WithOutput(o io.Writer) Option {
	return func(c *Config) {
		if o == nil {
			o = os.Stderr
		}
		c.Out = o
	}
}

// WithNoLock set config WithNoLock
func WithNoLock() Option {
	return func(c *Config) {
		c.WithNoLock = true
	}
}

// WithCaller set config WithCaller
func WithCaller() Option {
	return func(c *Config) {
		c.WithCaller = true
	}
}

// WithJSONEncoder set config Encoder with JSONEncoder
func WithJSONEncoder() Option {
	return func(c *Config) {
		c.Encoder = JSONEncoder
	}
}

// WithConsoleEncoder set config Encoder with ConsoleEncoder
func WithConsoleEncoder() Option {
	return func(c *Config) {
		c.Encoder = ConsoleEncoder
	}
}

// WithDebugLevel set config debug level
func WithDebugLevel() Option {
	return func(c *Config) {
		c.Level = DebugLevel
	}
}

// WithInfoLevel set config info level
func WithInfoLevel() Option {
	return func(c *Config) {
		c.Level = InfoLevel
	}
}

// WithWarnLevel set config warn level
func WithWarnLevel() Option {
	return func(c *Config) {
		c.Level = WarnLevel
	}
}

// WithErrorLevel set config error level
func WithErrorLevel() Option {
	return func(c *Config) {
		c.Level = ErrorLevel
	}
}

// WithPanicLevel set config panic level
func WithPanicLevel() Option {
	return func(c *Config) {
		c.Level = PanicLevel
	}
}

// WithFatalLevel set config fatal level
func WithFatalLevel() Option {
	return func(c *Config) {
		c.Level = FatalLevel
	}
}

// WithTimePattern set config TimePattern
func WithTimePattern(pattern string) Option {
	return func(c *Config) {
		c.TimePattern = pattern
	}
}
