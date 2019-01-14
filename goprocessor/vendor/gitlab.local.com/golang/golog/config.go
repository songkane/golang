// Package golog logger config
// Created by chenguolin 2018-12-26
package golog

import (
	"io"
	"os"
)

// Config for zap logger
type Config struct {
	level       LogLevel    //logger level
	encoder     EncoderType //json or console encoder
	withCaller  bool        //print the fileName & line number within the log
	out         io.Writer   //log out
	withNoLock  bool        //lock for the zap writer, default is false
	timePattern string      //time pattern
}

// Option set the option to Config
type Option func(c *Config)

// WithOutput set config Output
func WithOutput(o io.Writer) Option {
	return func(c *Config) {
		if o == nil {
			o = os.Stderr
		}
		c.out = o
	}
}

// withNoLock set config withNoLock
func WithNoLock() Option {
	return func(c *Config) {
		c.withNoLock = true
	}
}

// withCaller set config withCaller
func WithCaller() Option {
	return func(c *Config) {
		c.withCaller = true
	}
}

// WithJSONEncoder set config encoder with JSONEncoder
func WithJSONEncoder() Option {
	return func(c *Config) {
		c.encoder = JSONEncoder
	}
}

// WithConsoleEncoder set config encoder with ConsoleEncoder
func WithConsoleEncoder() Option {
	return func(c *Config) {
		c.encoder = ConsoleEncoder
	}
}

// WithDebugLevel set config debug level
func WithDebugLevel() Option {
	return func(c *Config) {
		c.level = DebugLevel
	}
}

// WithInfoLevel set config info level
func WithInfoLevel() Option {
	return func(c *Config) {
		c.level = InfoLevel
	}
}

// WithWarnLevel set config warn level
func WithWarnLevel() Option {
	return func(c *Config) {
		c.level = WarnLevel
	}
}

// WithErrorLevel set config error level
func WithErrorLevel() Option {
	return func(c *Config) {
		c.level = ErrorLevel
	}
}

// WithPanicLevel set config panic level
func WithPanicLevel() Option {
	return func(c *Config) {
		c.level = PanicLevel
	}
}

// WithFatalLevel set config fatal level
func WithFatalLevel() Option {
	return func(c *Config) {
		c.level = FatalLevel
	}
}

// WithTimePattern set config timePattern
func WithTimePattern(pattern string) Option {
	return func(c *Config) {
		if pattern == "" {
			pattern = DefaultTimePattern
		}
		c.timePattern = pattern
	}
}
