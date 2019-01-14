# gomysql
golang mysql 封装支出主从服务, 读取默认情况使用slave，写入默认情况使用master。
用户可以指定是否需要从master读取，调用`Mysql.Master()`函数即可保证从主库读取。

# sample
## open db
```
conf := &mysql.Config{}
conf.SetMaster("127.0.0.1")
conf.SetSlaves([]string{"127.0.0.1", "127.0.0.1"})
conf.SetPort(3306)
conf.SetUserName("root")
conf.SetPassword("root")
conf.SetDBName("test")
conf.SetMaxOpenConnCount(1000)
conf.SetMaxIdleConnCount(500)
conf.SetConnWaitTimeMs(5000)
conf.SetConnIdleTimeMs(21600000)
conf.SetConnTimeoutMs(5000)
conf.SetWriteTimeoutMs(5000)
conf.SetReadTimeoutMs(5000)

dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(err)
}
defer dbProxy.Close()
```

## create database
```
dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(fmt.Sprintf("NewMysql error:%s", err.Error()))
}
defer dbProxy.Close()

// create database
sql := string("create database sample_db")
res, err := dbProxy.Exec(sql)
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy.Exec error:%s", err.Error()))
}
fmt.Println(res)
```

## create table
```
dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(fmt.Sprintf("NewMysql error:%s", err.Error()))
}
defer dbProxy.Close()

// create table
sql := string("create table sample_table (id int, name varchar(255))")
res, err := dbProxy.Exec(sql)
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy.Exec error:%s", err.Error()))
}
fmt.Println(res)
```

## sql insert
```
dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(fmt.Sprintf("NewMysql error:%s", err.Error()))
}
defer dbProxy.Close()

args := []interface{}{
    time.Now().Unix(),
"cgl",
}

res, err := dbProxy.Exec("insert into sample_table (id, name) values (?, ?)", args...)
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy.Exec insert error:%s", err.Error()))
}
fmt.Println(res.LastInsertId())
fmt.Println(res.RowsAffected())
```

## sql select
```
dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(fmt.Sprintf("NewMysql error:%s", err.Error()))
}
defer dbProxy.Close()

rows, err := dbProxy.Query("select * from sample_table")
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy.Query select error:%s", err.Error()))
}

// scan
defer rows.Close()
for rows.Next() {
    var id int
	var name string
	err := rows.Scan(&id, &name)
	if err != nil {
	    fmt.Println(fmt.Sprintf("rows.Scan error:%s", err.Error()))
	}
	fmt.Println(id, name)
}
```

## sql update
```
dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(fmt.Sprintf("NewMysql error:%s", err.Error()))
}
defer dbProxy.Close()

// update
res, err := dbProxy.Exec("update sample_table set id = 1 where name = 'cgl'")
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy.Exec update error:%s", err.Error()))
}
fmt.Println(res.RowsAffected())
fmt.Println(res.LastInsertId())
```

## sql delete
```
dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(fmt.Sprintf("NewMysql error:%s", err.Error()))
}
defer dbProxy.Close()

// delete
res, err := dbProxy.Exec("delete from sample_table where id = 1")
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy.Exec delete error:%s", err.Error()))
}
fmt.Println(res.RowsAffected())
fmt.Println(res.LastInsertId())
```

## sql tx
```
dbProxy, err := mysql.NewMysql(conf)
if err != nil {
    fmt.Println(fmt.Sprintf("NewMysql error:%s", err.Error()))
}
defer dbProxy.Close()

// begin
tx, err := dbProxy.Begin()
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy begin error:%s", err.Error()))
}

// sql
args := []interface{}{
    time.Now().Unix(),
	"cgl",
}

res, err := dbProxy.Exec("insert into sample_table (id, name) values (?, ?)", args...)
if err != nil {
    fmt.Println(fmt.Sprintf("dbProxy.Exec insert error:%s", err.Error()))
}
fmt.Println(res.RowsAffected())
fmt.Println(res.LastInsertId())

// commit
err = tx.Commit()

// err rollback
if err != nil {
    tx.Rollback()
}
```