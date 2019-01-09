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
// max_open_conns = 100 (设置最大打开的连接数,默认值为0表示不限制)
// max_idle_conns = 50 (设置连接池数量)
// conn_wait_time_ms = 200
// conn_idle_time_ms = 21600000
// conn_timeout_ms = 5000
// read_timeout_ms = 3000
// write_timeout_ms = 3000
type Config struct {
	Master         string   `json:"master"`            //required mysql master
	Slaves         []string `json:"slaves"`            //required mysql slaves
	Port           int      `json:"port"`              //required mysql port
	UserName       string   `json:"user_name"`         //required mysql user name
	Password       string   `json:"password"`          //required mysql password
	DBName         string   `json:"db_name"`           //required mysql database name
	MaxOpenConns   int      `json:"max_open_conns"`    //optional mysql max open connections count
	MaxIdleConns   int      `json:"max_idle_conns"`    //optional mysql max idle connections count
	ConnWaitTimeMs int      `json:"conn_wait_time_ms"` //optional mysql connection wait time ms
	ConnIdleTimeMs int      `json:"conn_idle_time_ms"` //optional mysql connection idle time ms
	ConnTimeoutMs  int      `json:"conn_timeout_ms"`   //optional mysql connection timeout ms
	ReadTimeoutMs  int      `json:"read_timeout_ms"`   //optional mysql read timeout ms
	WriteTimeoutMs int      `json:"write_timeout_ms"`  //optional mysql write timeout ms
}

// validate config config
// must have required fields
func (c *Config) validate() bool {
	if len(c.Master) <= 0 || len(c.Slaves) <= 0 || c.Port == 0 ||
		len(c.UserName) <= 0 || len(c.Password) <= 0 || len(c.DBName) <= 0 {
		return false
	}

	return true
}

// SetMaster set mysql master server
func (c *Config) SetMaster(master string) {
	c.Master = master
}

// SetSlaves set mysql slaves
func (c *Config) SetSlaves(slaves []string) {
	c.Slaves = slaves
}

// SetPort set mysql server port
func (c *Config) SetPort(port int) {
	c.Port = port
}

// SetUserName set mysql user name
func (c *Config) SetUserName(userName string) {
	c.UserName = userName
}

// SetPassword set mysql password
func (c *Config) SetPassword(password string) {
	c.Password = password
}

// SetDBName set mysql database
func (c *Config) SetDBName(dbName string) {
	c.DBName = dbName
}

// SetMaxOpenConns set mysql max open connections
// 设置最大的连接数, 可以避免并发太高导致连接mysql出现too many connections的错误。
func (c *Config) SetMaxOpenConns(maxOpenConns int) {
	c.MaxOpenConns = maxOpenConns
}

// SetMaxIdleConns set mysql max idle connections
// 设置连接池数量，通常设置小于最大的连接数
func (c *Config) SetMaxIdleConns(maxIdleConns int) {
	c.MaxIdleConns = maxIdleConns
}

// SetMaxWaitTimeMs set max wait time ms
func (c *Config) SetConnWaitTimeMs(connWaitTimeMs int) {
	c.ConnWaitTimeMs = connWaitTimeMs
}

// SetMaxIdleTimeMs set max idle time ms
func (c *Config) SetConnIdleTimeMs(connIdleTimeMs int) {
	c.ConnIdleTimeMs = connIdleTimeMs
}

// SetConnTimeoutMs set connection timeout ms
func (c *Config) SetConnTimeoutMs(connTimeoutMs int) {
	c.ConnTimeoutMs = connTimeoutMs
}

// SetReadTimeoutMs set read timeout ms
func (c *Config) SetReadTimeoutMs(readTimeoutMs int) {
	c.ReadTimeoutMs = readTimeoutMs
}

// SetWriteTimeoutMs set write timeout ms
func (c *Config) SetWriteTimeoutMs(writeTimeoutMs int) {
	c.WriteTimeoutMs = writeTimeoutMs
}
