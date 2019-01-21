// Package cron handle
// Created by chenguolin 2018-01-04
package cron

// handle interface
// All handle need implement DoProcess function
type handle interface {
	DoProcess()
}
