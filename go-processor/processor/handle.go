// Package processor Hanlde define
// Created by chenguolin 2019-01-13
package processor

// Handle interface
type Handle interface {
	Process(record Record)
}
