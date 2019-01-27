// Package mysql wrapper
// Created by chenguolin 2019-01-06
package mysql

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gitlab.local.com/golang/go-mysql/sql"
)

// Error define
var (
	// ErrInvalidConfig invalid mysql config
	ErrInvalidConfig = errors.New("mysql config invalid")
)

// Mysql define
type Mysql struct {
	master *sql.DB   //mysql master server
	slaves []*sql.DB //mysql slaves server
}

// NewMysql new mysql
func NewMysql(conf *Config) (*Mysql, error) {
	// check conf
	if conf == nil || !conf.validate() {
		return nil, ErrInvalidConfig
	}

	// new master DB
	masterDB, err := newMasterDB(conf)
	if err != nil {
		return nil, err
	}

	// new slaves DB
	slavesDB, err := newSlavesDB(conf)
	if err != nil {
		return nil, err
	}

	// new mysql
	mysql := &Mysql{
		master: masterDB,
		slaves: slavesDB,
	}

	return mysql, nil
}

// generateMasterDB generate master DB
func newMasterDB(conf *Config) (*sql.DB, error) {
	// generate master DSN
	masterDsn := generateMasterDSN(conf)

	// open DB
	// master DB
	master, err := sql.Open(DefaultDriver, masterDsn)
	if err != nil {
		return nil, err
	}

	// set DB parameters
	setDBParameters(master, conf)

	return master, nil
}

// generateMasterDB generate master DB
func newSlavesDB(conf *Config) ([]*sql.DB, error) {
	// generate slaves DSN
	slavesDsn := generateSlavesDSN(conf)

	// slaves DB
	slaves := make([]*sql.DB, 0)
	for _, slaveDsn := range slavesDsn {
		db, err := sql.Open(DefaultDriver, slaveDsn)
		if err != nil {
			return nil, err
		}

		// set DB parameters
		setDBParameters(db, conf)
		slaves = append(slaves, db)
	}

	return slaves, nil
}

// setDBParameters
// @SetMaxOpenConnCount: sets the maximum number of open connections to the database. The default is 0 (unlimited).
// @SetMaxIdleConnCount: sets the maximum number of connections in the idle connection pool. If n <= 0, no idle connections are retained.
// @SetConnWaitTimeout: sets the maximum amount of time a connRequest may wait to be satisfied. Default ConnWaitTimeout is 3 seconds
// @SetConnIdleTimeout: sets the maximum time that an idle connection can remain idle in the pool. Default idle connections will not be removed.
func setDBParameters(db *sql.DB, conf *Config) {
	// SetMaxOpenConnCount
	if conf.maxOpenConnCount > 0 {
		db.SetMaxOpenConnCount(conf.maxOpenConnCount)
	}

	// SetMaxIdleConnCount
	if conf.maxIdleConnCount > 0 {
		db.SetMaxIdleConnCount(conf.maxIdleConnCount)
	}

	// SetConnWaitTimeout
	if conf.connWaitTimeMs > 0 {
		db.SetConnWaitTimeout(time.Duration(conf.connWaitTimeMs) * time.Millisecond)
	}

	// SetConnIdleTimeout
	if conf.connIdleTimeMs > 0 {
		db.SetConnIdleTimeout(time.Duration(conf.connIdleTimeMs) * time.Millisecond)
	}
}

// generateMasterDSN generate master data source name
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
// tcp protocol: [tcp://addr/]dbname/user/password[?params]
// unix sock protocol：[unix://sockpath/]dbname/user/password[?params]
func generateMasterDSN(conf *Config) string {
	// format dsn default use TCP protocol
	masterDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.userName, conf.password,
		conf.master, conf.port, conf.dbName)
	masterDsn = setDsnTimeoutParameters(masterDsn, conf)

	return masterDsn
}

// generateSlavesDSN generate slaves data source name
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
// tcp protocol: [tcp://addr/]dbname/user/password[?params]
// unix sock protocol：[unix://sockpath/]dbname/user/password[?params]
func generateSlavesDSN(conf *Config) []string {
	// format dsn default use TCP protocol
	slavesDsn := make([]string, 0)
	for _, slave := range conf.slaves {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.userName, conf.password,
			slave, conf.port, conf.dbName)
		dsn = setDsnTimeoutParameters(dsn, conf)
		slavesDsn = append(slavesDsn, dsn)
	}

	return slavesDsn
}

// setDsnTimeoutParameters
// @timeout: Timeout for establishing connections, aka dial timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
// @readTimeout: I/O read timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
// @writeTimeout: I/O write timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
func setDsnTimeoutParameters(dsn string, conf *Config) string {
	// ConnTimeout
	if conf.connTimeoutMs > 0 {
		dsn = fmt.Sprintf("%s?timeout=%dms", dsn, conf.connTimeoutMs)
	}
	// ReadTimeout
	if conf.readTimeoutMs > 0 {
		dsn = fmt.Sprintf("%s&readTimeout=%dms", dsn, conf.readTimeoutMs)
	}
	// WriteTimeout
	if conf.writeTimeoutMs > 0 {
		dsn = fmt.Sprintf("%s&writeTimeout=%dms", dsn, conf.writeTimeoutMs)
	}

	return dsn
}

// master return master DB
func (m *Mysql) Master() *sql.DB {
	return m.master
}

// Conn returns a single connection by either opening a new connection
// or returning an existing connection from the connection pool. Conn will
// block until either a connection is returned or ctx is canceled.
// Queries run on the same Conn will be run in the same database session.
//
// Every Conn must be returned to the database pool after use by
// calling Conn.Close.
func (m *Mysql) Conn(ctx context.Context) (*sql.Conn, error) {
	// Default use slave server
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(m.slaves))
	db := m.slaves[pos]
	return db.Conn(ctx)
}

// Close closes the database, releasing any open resources.
// It is rare to Close a DB, as the DB handle is meant to be
// long-lived and shared between many goroutines.
func (m *Mysql) Close() error {
	// close master
	err := m.master.Close()
	if err != nil {
		return err
	}

	// close slaves
	for _, slave := range m.slaves {
		err := slave.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (m *Mysql) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// Default use slave server
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(m.slaves))
	db := m.slaves[pos]
	return db.Query(query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (m *Mysql) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// Default use slave server
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(m.slaves))
	db := m.slaves[pos]
	return db.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (m *Mysql) QueryRow(query string, args ...interface{}) *sql.Row {
	// Default use slave server
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(m.slaves))
	db := m.slaves[pos]
	return db.QueryRow(query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (m *Mysql) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// Default use slave server
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(m.slaves))
	db := m.slaves[pos]
	return db.QueryRowContext(ctx, query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (m *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	// Default use master server
	db := m.master
	// prepare
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	// exec
	return stmt.Exec(args...)
}

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (m *Mysql) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Default use master server
	db := m.master
	// prepare
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	// exec
	return stmt.ExecContext(ctx, args...)
}

// Begin starts a transaction. The default isolation level is dependent on
// the driver.
func (m *Mysql) Begin() (*sql.Tx, error) {
	// Default use slave server
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(m.slaves))
	db := m.slaves[pos]
	return db.Begin()
}

// BeginTx starts a transaction.
//
// The provided context is used until the transaction is committed or rolled back.
// If the context is canceled, the sql package will roll back
// the transaction. Tx.Commit will return an error if the context provided to
// BeginTx is canceled.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
func (m *Mysql) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	// Default use slave server
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(m.slaves))
	db := m.slaves[pos]
	return db.BeginTx(ctx, opts)
}
