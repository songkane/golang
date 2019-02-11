// Package sql register mysql driver
// Created by chenguolin 2019-01-07
package sql

import (
	"github.com/go-sql-driver/mysql"
)

const (
	// DefaultDriver default mysql driver
	DefaultDriver = "mysql"
)

func init() {
	Register(DefaultDriver, &mysql.MySQLDriver{})
}
