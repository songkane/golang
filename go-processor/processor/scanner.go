// Package processor processor define
// Created by chenguolin 2019-01-13
package processor

// Record define
type Record interface{}

// Scanner interface
type Scanner interface {
	// Start scanner
	Start()
	// Stop scanner
	Stop()
	// Next record
	Next() (Record, bool)
	// IsStopped check scanner is stop
	IsStopped() bool
}
