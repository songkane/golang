// Package mysql init
// Created by chenguolin 2019-01-07
package mysql

import (
	"github.com/go-sql-driver/mysql"
	"gitlab.local.com/golang/go-mysql/sql"
)

const (
	// DefaultDriver default mysql driver
	DefaultDriver = "mysql"
)

func init() {
	sql.Register(DefaultDriver, &mysql.MySQLDriver{})
}
