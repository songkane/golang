// Package user UserRepo
// Created by chenguolin 2018-05-02
package user

import (
	"fmt"
	"strings"
	"time"

	"gitlab.local.com/golang/go-http/bizerror"
	"gitlab.local.com/golang/go-mysql"
	"gitlab.local.com/golang/go-mysql/sql"
)

// repo 用户表操作Repo对象
type repo struct {
	db           *mysql.Mysql
	tableName    string
	allFields    string
	insertFields string
}

// newRepo 实例化一个UserRepo
func newRepo(db *mysql.Mysql) *repo {
	if db == nil {
		panic("[UsersRepo - newRepo] panic")
	}

	return &repo{
		db:           db,
		tableName:    "users",
		allFields:    "id, uid, name, phone, create_time, update_time",
		insertFields: "uid, name, phone, create_time, update_time",
	}
}

// insertUser 新增用户
func (ur *repo) insertUser(user *userTableModel) error {
	// 1. 插入mysql
	fieldsCount := len(strings.Split(ur.insertFields, ","))
	sqlString := fmt.Sprintf("insert into %s (%s) values (%s)", ur.tableName, ur.insertFields,
		strings.TrimRight(strings.Repeat("?,", fieldsCount), ","))
	values := []interface{}{
		user.UID,
		user.Name,
		user.Phone,
		user.CreateTime,
		user.UpdateTime,
	}

	result, err := ur.db.Exec(sqlString, values...)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

// deleteUser 删除用户
func (ur *repo) deleteUser(uid int64) error {
	// TODO
	return nil
}

// updateUser 更新用户信息
func (ur *repo) updateUser(uid int64, newName, newPhone string) error {
	updateTime := time.Now().Unix() //10位时间戳
	sqlString := fmt.Sprintf("update %s set name=?, phone=?, update_time=? where uid=?", ur.tableName)
	values := []interface{}{
		newName,
		newPhone,
		updateTime,
	}

	result, err := ur.db.Exec(sqlString, values...)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// selectUser 查询用户信息
func (ur *repo) selectUser(uid int64) (*userTableModel, error) {
	sqlString := fmt.Sprintf("select %s from %s where `uid` = ?",
		ur.allFields, ur.tableName)

	rows, err := ur.db.Query(sqlString, uid)
	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, bizerror.ErrMysqlRecordNotFound
	}

	userModelList, err := ur.convert(rows)
	if err != nil {
		return nil, err
	}

	if len(userModelList) <= 0 {
		return nil, bizerror.ErrMysqlRecordNotFound
	}

	return userModelList[0], nil
}

// convert sql row转换成UserModel
func (ur *repo) convert(rows *sql.Rows) ([]*userTableModel, error) {
	defer rows.Close()
	userModelList := make([]*userTableModel, 0)

	for rows.Next() {
		record := userTableModel{}
		err := rows.Scan(&record.ID, &record.UID, &record.Name,
			&record.Phone, &record.CreateTime, &record.UpdateTime)
		if err != nil {
			return nil, err
		}
		userModelList = append(userModelList, &record)
	}

	return userModelList, nil
}
