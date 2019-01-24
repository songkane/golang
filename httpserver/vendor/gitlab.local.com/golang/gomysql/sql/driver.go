// Package sql driver
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Type conversions for Scan.
package sql

import (
	"database/sql/driver"
	"errors"
)

// ErrRemoveArgument may be returned from NamedValueChecker to instruct the
// sql package to not pass the argument to the driver query interface.
// Return when accepting query specific options or structures that aren't
// SQL query arguments.
var ErrRemoveArgument = errors.New("driver: remove argument from query")

// NamedValueChecker may be optionally implemented by Conn or Stmt. It provides
// the driver more control to handle Go and database types beyond the default
// Values types allowed.
//
// The sql package checks for value checkers in the following order,
// stopping at the first found match: Stmt.NamedValueChecker, Conn.NamedValueChecker,
// Stmt.ColumnConverter, DefaultParameterConverter.
//
// If CheckNamedValue returns ErrRemoveArgument, the NamedValue will not be included in
// the final query arguments. This may be used to pass special options to
// the query itself.
//
// If ErrSkip is returned the column converter error checking
// path is used for the argument. Drivers may wish to return ErrSkip after
// they have exhausted their own special cases.
type NamedValueChecker interface {
	// CheckNamedValue is called before passing arguments to the driver
	// and is called in place of any ColumnConverter. CheckNamedValue must do type
	// validation and conversion as appropriate for the driver.
	CheckNamedValue(*driver.NamedValue) error
}
