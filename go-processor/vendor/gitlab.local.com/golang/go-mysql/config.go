// Package mysql config
// Created by chenguolin 2019-01-06
package mysql

// Config mysql config
// master = "127.0.0.1"
// slaves = ["127.0.0.1", "127.0.0.2"]
// port = 3306
// username = "root"
// password = "root"
// dbname = "test"
// max_open_conn_count = 100 (设置最大打开的连接数,默认值为0表示不限制)
// max_idle_conn_count = 50 (设置连接池数量)
// conn_wait_time_ms = 200
// conn_idle_time_ms = 21600000
// conn_timeout_ms = 5000 (dsn parameter)
// read_timeout_ms = 3000 (dsn parameter)
// write_timeout_ms = 3000 (dsn parameter)
type Config struct {
	master           string   //required mysql master
	slaves           []string //required mysql slaves
	port             int      //required mysql port
	userName         string   //required mysql user name
	password         string   //required mysql password
	dbName           string   //required mysql database name
	maxOpenConnCount int      //optional mysql max open connections count
	maxIdleConnCount int      //optional mysql max idle connections count
	connWaitTimeMs   int      //optional mysql connection wait time ms
	connIdleTimeMs   int      //optional mysql connection idle time ms
	connTimeoutMs    int      //optional mysql connection timeout ms
	readTimeoutMs    int      //optional mysql read timeout ms
	writeTimeoutMs   int      //optional mysql write timeout ms
}

// validate config config
// must have required fields
func (c *Config) validate() bool {
	if len(c.master) <= 0 || len(c.slaves) <= 0 || c.port == 0 ||
		len(c.userName) <= 0 || len(c.password) <= 0 || len(c.dbName) <= 0 {
		return false
	}

	return true
}

// SetMaster set mysql master server
func (c *Config) SetMaster(master string) {
	c.master = master
}

// SetSlaves set mysql slaves
func (c *Config) SetSlaves(slaves []string) {
	c.slaves = slaves
}

// SetPort set mysql server port
func (c *Config) SetPort(port int) {
	c.port = port
}

// SetUserName set mysql user name
func (c *Config) SetUserName(userName string) {
	c.userName = userName
}

// SetPassword set mysql password
func (c *Config) SetPassword(password string) {
	c.password = password
}

// SetDBName set mysql database
func (c *Config) SetDBName(dbName string) {
	c.dbName = dbName
}

// SetMaxOpenConnCount set mysql max open connections
// 设置最大的连接数, 可以避免并发太高导致连接mysql出现too many connections的错误。
func (c *Config) SetMaxOpenConnCount(maxOpenConnCount int) {
	c.maxOpenConnCount = maxOpenConnCount
}

// SetMaxIdleConnCount set mysql max idle connections
// 设置连接池数量，通常设置小于最大的连接数
func (c *Config) SetMaxIdleConnCount(maxIdleConnCount int) {
	c.maxIdleConnCount = maxIdleConnCount
}

// SetConnWaitTimeMs set max wait time ms
func (c *Config) SetConnWaitTimeMs(connWaitTimeMs int) {
	c.connWaitTimeMs = connWaitTimeMs
}

// SetConnIdleTimeMs set max idle time ms
func (c *Config) SetConnIdleTimeMs(connIdleTimeMs int) {
	c.connIdleTimeMs = connIdleTimeMs
}

// SetConnTimeoutMs set connection timeout ms
func (c *Config) SetConnTimeoutMs(connTimeoutMs int) {
	c.connTimeoutMs = connTimeoutMs
}

// SetReadTimeoutMs set read timeout ms
func (c *Config) SetReadTimeoutMs(readTimeoutMs int) {
	c.readTimeoutMs = readTimeoutMs
}

// SetWriteTimeoutMs set write timeout ms
func (c *Config) SetWriteTimeoutMs(writeTimeoutMs int) {
	c.writeTimeoutMs = writeTimeoutMs
}
