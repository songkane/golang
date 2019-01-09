// Package mysql config unit test
// Created by chenguolin 2019-01-08
package mysql

import "testing"

func TestConfig_validate(t *testing.T) {
	config := &Config{}
	if config.validate() {
		t.Fatal("TestConfig_validate error")
	}

	// case 1
	config.SetMaster("127.0.0.1")
	if config.validate() {
		t.Fatal("TestConfig_validate case 1 error")
	}
	// case 2
	slaves := []string{"127.0.0.1", "127.0.0.2"}
	config.SetSlaves(slaves)
	if config.validate() {
		t.Fatal("TestConfig_validate case 2 error")
	}
	// case 3
	config.SetPort(3306)
	if config.validate() {
		t.Fatal("TestConfig_validate case 3 error")
	}
	// case 4
	config.SetUserName("root")
	if config.validate() {
		t.Fatal("TestConfig_validate case 4 error")
	}
	// case 5
	config.SetPassword("root")
	if config.validate() {
		t.Fatal("TestConfig_validate case 5 error")
	}
	// case 6
	config.SetDBName("test")
	if !config.validate() {
		t.Fatal("TestConfig_validate case 6 error")
	}
	// case 7
	config.SetMaxOpenConns(100)
	if !config.validate() {
		t.Fatal("TestConfig_validate case 7 error")
	}
	// case 8
	config.SetMaxIdleConns(50)
	if !config.validate() {
		t.Fatal("TestConfig_validate case 8 error")
	}
}

func TestConfig_SetMaster(t *testing.T) {
	config := &Config{}
	if config.Master != "" {
		t.Fatal("TestConfig_SetMaster != \"\"")
	}

	config.SetMaster("127.0.0.1")
	if config.Master != "127.0.0.1" {
		t.Fatal("TestConfig_SetMaster !=\"127.0.0.1\"")
	}
}

func TestConfig_SetSlaves(t *testing.T) {
	config := &Config{}
	if config.Slaves != nil {
		t.Fatal("TestConfig_SetSlaves != nil")
	}

	slaves := []string{"127.0.0.1", "127.0.0.2"}
	config.SetSlaves(slaves)
	if len(config.Slaves) != len(slaves) {
		t.Fatal("TestConfig_SetSlaves != len(slaves)")
	}
}

func TestConfig_SetPort(t *testing.T) {
	config := &Config{}
	if config.Port != 0 {
		t.Fatal("TestConfig_SetPort != 0")
	}

	config.SetPort(3306)
	if config.Port != 3306 {
		t.Fatal("TestConfig_SetPort != 3306")
	}
}

func TestConfig_SetUserName(t *testing.T) {
	config := &Config{}
	if config.UserName != "" {
		t.Fatal("TestConfig_SetUserName != \"\"")
	}

	config.SetUserName("root")
	if config.UserName != "root" {
		t.Fatal("TestConfig_SetUserName error")
	}
}

func TestConfig_SetPassword(t *testing.T) {
	config := &Config{}
	if config.Password != "" {
		t.Fatal("TestConfig_SetPassword != \"\"")
	}

	config.SetPassword("root")
	if config.Password != "root" {
		t.Fatal("TestConfig_SetPassword error")
	}
}

func TestConfig_SetDBName(t *testing.T) {
	config := &Config{}
	if config.DBName != "" {
		t.Fatal("TestConfig_SetDBName != \"\"")
	}

	config.SetDBName("test")
	if config.DBName != "test" {
		t.Fatal("TestConfig_SetDBName error")
	}
}

func TestConfig_SetMaxOpenConns(t *testing.T) {
	config := &Config{}
	if config.MaxOpenConns != 0 {
		t.Fatal("TestConfig_SetMaxOpenConns != 0")
	}

	config.SetMaxOpenConns(100)
	if config.MaxOpenConns != 100 {
		t.Fatal("TestConfig_SetMaxOpenConns != 100")
	}
}

func TestConfig_SetMaxIdleConns(t *testing.T) {
	config := &Config{}
	if config.MaxIdleConns != 0 {
		t.Fatal("TestConfig_SetMaxIdleConns != 0")
	}

	config.SetMaxIdleConns(50)
	if config.MaxIdleConns != 50 {
		t.Fatal("TestConfig_SetMaxIdleConns != 50")
	}
}

func TestConfig_SetConnTimeoutMs(t *testing.T) {
	config := &Config{}
	if config.ConnTimeoutMs != 0 {
		t.Fatal("TestConfig_SetConnTimeoutMs != 0")
	}

	config.SetConnTimeoutMs(1000)
	if config.ConnTimeoutMs != 1000 {
		t.Fatal("TestConfig_SetConnTimeoutMs != 1000")
	}
}

func TestConfig_SetReadTimeoutMs(t *testing.T) {
	config := &Config{}
	if config.ReadTimeoutMs != 0 {
		t.Fatal("TestConfig_SetReadTimeoutMs != 0")
	}

	config.SetReadTimeoutMs(1000)
	if config.ReadTimeoutMs != 1000 {
		t.Fatal("TestConfig_SetReadTimeoutMs != 1000")
	}
}

func TestConfig_SetWriteTimeoutMs(t *testing.T) {
	config := &Config{}
	if config.WriteTimeoutMs != 0 {
		t.Fatal("TestConfig_SetWriteTimeoutMs != 0")
	}

	config.SetWriteTimeoutMs(1000)
	if config.WriteTimeoutMs != 1000 {
		t.Fatal("TestConfig_SetWriteTimeoutMs != 1000")
	}
}

func TestConfig_SetConnWaitTimeMs(t *testing.T) {
	config := &Config{}
	if config.ConnWaitTimeMs != 0 {
		t.Fatal("TestConfig_SetMaxWaitTimeMs != 0")
	}

	config.SetConnWaitTimeMs(1000)
	if config.ConnWaitTimeMs != 1000 {
		t.Fatal("TestConfig_SetMaxWaitTimeMs != 1000")
	}
}

func TestConfig_SetConnIdleTimeMs(t *testing.T) {
	config := &Config{}
	if config.ConnIdleTimeMs != 0 {
		t.Fatal("TestConfig_SetMaxIdleTimeMs != 0")
	}

	config.SetConnIdleTimeMs(1000)
	if config.ConnIdleTimeMs != 1000 {
		t.Fatal("TestConfig_SetMaxIdleTimeMs != 1000")
	}
}
