// Package mysql unit test
// Created by chenguolin
package mysql

import (
	"context"
	"testing"

	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestMysql_NewMysql(t *testing.T) {
	// case 1
	mysql, err := NewMysql(nil)
	if mysql != nil {
		t.Fatal("TestMysql_NewMysql case 1 mysql != nil")
	}
	if err == nil {
		t.Fatal("TestMysql_NewMysql case 1 err == nil")
	}

	// case 2
	conf := &Config{}
	mysql, err = NewMysql(conf)
	if mysql != nil {
		t.Fatal("TestMysql_NewMysql case 2 mysql != nil")
	}
	if err == nil {
		t.Fatal("TestMysql_NewMysql case 2 err == nil")
	}

	// case 3
	// normal case
	conf = &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	mysql, err = NewMysql(conf)
	if mysql == nil {
		t.Fatal("TestMysql_NewMysql case 3 mysql == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_NewMysql case 3 err != nil")
	}

	// case 4
	// invalid username and password
	conf = &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root2")
	conf.SetPassword("root2")
	conf.SetDBName("test")
	mysql, err = NewMysql(conf)
	if mysql == nil {
		t.Fatal("TestMysql_NewMysql case 4 mysql == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_NewMysql case 4 err != nil")
	}

	// case 5
	conf = &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	mysql, err = NewMysql(conf)
	if mysql == nil {
		t.Fatal("TestMysql_NewMysql case 5 mysql == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_NewMysql case 5 err != nil")
	}
}

func TestMysql_newMasterDB(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	db, err := newMasterDB(conf)
	if db == nil {
		t.Fatal("TestMysql_newMasterDB db == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_newMasterDB err != nil")
	}
}

func TestMysql_newSlavesDB(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	db, err := newSlavesDB(conf)
	if db == nil {
		t.Fatal("TestMysql_newSlavesDB db == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_newSlavesDB err != nil")
	}
}

func TestMysql_setDBParameters(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	db, _ := newMasterDB(conf)
	setDBParameters(db, conf)
}

func TestMysql_generateMasterDSN(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")

	dsn := generateMasterDSN(conf)
	if dsn != string("root:root@tcp(127.0.0.1:3306)/test") {
		t.Fatal("TestMysql_generateMasterDSN dsn != string(\"root:root@tcp(127.0.0.1:3306)/test\")")
	}
}

func TestMysql_generateSlavesDSN(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")

	dsns := generateSlavesDSN(conf)
	if len(dsns) != 2 {
		t.Fatal("TestMysql_generateSlavesDSN len(dsns) != 2")
	}
	if dsns[0] != string("root:root@tcp(127.0.0.1:3306)/test") {
		t.Fatal("TestMysql_generateSlavesDSN dsn[0] != string(\"root:root@tcp(127.0.0.1:3306)/test\")")
	}
	if dsns[1] != string("root:root@tcp(127.0.0.1:3306)/test") {
		t.Fatal("TestMysql_generateSlavesDSN dsn[1] != string(\"root:root@tcp(127.0.0.1:3306)/test\")")
	}
}

func TestMysql_setDsnTimeoutParameters(t *testing.T) {
	// case 1
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")

	dsn := generateMasterDSN(conf)
	dsn = setDsnTimeoutParameters(dsn, conf)
	if dsn != string("root:root@tcp(127.0.0.1:3306)/test") {
		t.Fatal("TestMysql_setDsnTimeoutParameters dsn != string(\"root:root@tcp(127.0.0.1:3306)/test\")")
	}

	// case 2
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)
	dsn = setDsnTimeoutParameters(dsn, conf)
	if dsn != string("root:root@tcp(127.0.0.1:3306)/test?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms") {
		t.Fatal("TestMysql_setDsnTimeoutParameters dsn != string(\"root:root@tcp(127.0.0.1:3306)/test?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms\")")
	}
}

func TestMysql_Master(t *testing.T) {
	// case 1
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	if mysql.Master() == nil {
		t.Fatal("TestMysql_Master mysql.Master() == nil")
	}

	// 本地数据库必须要有userinfo表
	rows, err := mysql.Master().Query("select * from userinfo")
	if err != nil {
		t.Fatal("TestMysql_Query err != nil")
	}
	if rows == nil {
		t.Fatal("TestMysql_Query rows == nil")
	}
}

func TestMysql_Conn(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	conn, err := mysql.Conn(context.Background())
	if conn == nil {
		t.Fatal("TestMysql_Conn conn == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_Conn err != nil")
	}
	defer conn.Close()
}

func TestMysql_Close(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	err := mysql.Close()
	if err != nil {
		t.Fatal("TestMysql_Close err != nil")
	}
}

func TestMysql_Query(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	// 本地数据库必须要有userinfo表
	rows, err := mysql.Query("select * from `userinfo`")
	fmt.Println(err)
	if err != nil {
		t.Fatal("TestMysql_Query err != nil")
	}
	if rows == nil {
		t.Fatal("TestMysql_Query rows == nil")
	}

	fmt.Println(rows)
}

func TestMysql_QueryContext(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	// 本地数据库必须要有userinfo表
	rows, err := mysql.QueryContext(context.Background(), "select * from `userinfo`")
	fmt.Println(err)
	if err != nil {
		t.Fatal("TestMysql_QueryContext err != nil")
	}
	if rows == nil {
		t.Fatal("TestMysql_QueryContext rows == nil")
	}

	fmt.Println(rows)
}

func TestMysql_QueryRow(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	// 本地数据库必须要有userinfo表
	rows := mysql.QueryRow("select * from `userinfo` limit 1")
	fmt.Println(rows)
	if rows == nil {
		t.Fatal("TestMysql_QueryRow rows == nil")
	}
}

func TestMysql_QueryRowContext(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	// 本地数据库必须要有userinfo表
	rows := mysql.QueryRowContext(context.Background(), "select * from `userinfo` limit 1")
	fmt.Println(rows)
	if rows == nil {
		t.Fatal("TestMysql_QueryRowContext rows == nil")
	}
}

func TestMysql_Exec(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	// 本地数据库必须要有userinfo表
	args := []interface{}{
		time.Now().Unix(),
		"cgl2",
		"test2",
		"2019-01-09",
	}

	res, err := mysql.Exec("insert into userinfo (uid, username, department, created) values (?, ?, ?, ?)", args...)
	if err != nil {
		t.Fatal("TestMysql_Exec err != nil")
	}
	if res == nil {
		t.Fatal("TestMysql_Exec res == nil")
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}

func TestMysql_ExecContext(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	// 本地数据库必须要有userinfo表
	args := []interface{}{
		time.Now().Unix(),
		"cgl2",
		"test2",
		"2019-01-09",
	}

	res, err := mysql.ExecContext(context.Background(), "insert into userinfo (uid, username, department, created) values (?, ?, ?, ?)", args...)
	if err != nil {
		t.Fatal("TestMysql_ExecContext err != nil")
	}
	if res == nil {
		t.Fatal("TestMysql_ExecContext res == nil")
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}

func TestMysql_Begin(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	tx, err := mysql.Begin()
	if tx == nil {
		t.Fatal("TestMysql_Begin tx == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_Begin err != nil")
	}

	// 本地数据库必须要有userinfo表
	args := []interface{}{
		time.Now().Unix(),
		"cgl2",
		"test2",
		"2019-01-09",
	}
	res, err := tx.Exec("insert into userinfo (uid, username, department, created) values (?, ?, ?, ?)", args...)
	if err != nil {
		t.Fatal("TestMysql_Begin err != nil")
	}
	if res == nil {
		t.Fatal("TestMysql_Begin res == nil")
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	// commit
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			t.Fatal("TestMysql_Begin tx.Rollback err != nil")
		}
	}
}

func TestMysql_BeginTx(t *testing.T) {
	conf := &Config{}
	conf.SetMaster("127.0.0.1")
	conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
	conf.SetPort(3306)
	conf.SetUserName("root")
	conf.SetPassword("root")
	conf.SetDBName("test")
	conf.SetMaxOpenConns(1000)
	conf.SetMaxIdleConns(500)
	conf.SetConnWaitTimeMs(5000)
	conf.SetConnIdleTimeMs(21600000)
	conf.SetConnTimeoutMs(5000)
	conf.SetWriteTimeoutMs(5000)
	conf.SetReadTimeoutMs(5000)

	mysql, _ := NewMysql(conf)
	tx, err := mysql.BeginTx(context.Background(), nil)
	if tx == nil {
		t.Fatal("TestMysql_Begin tx == nil")
	}
	if err != nil {
		t.Fatal("TestMysql_Begin err != nil")
	}

	// 本地数据库必须要有userinfo表
	args := []interface{}{
		time.Now().Unix(),
		"cgl2",
		"test2",
		"2019-01-09",
	}
	res, err := tx.Exec("insert into userinfo (uid, username, department, created) values (?, ?, ?, ?)", args...)
	if err != nil {
		t.Fatal("TestMysql_Begin err != nil")
	}
	if res == nil {
		t.Fatal("TestMysql_Begin res == nil")
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	// commit
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			t.Fatal("TestMysql_Begin tx.Rollback err != nil")
		}
	}
}
