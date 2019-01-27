// Package user mysql table schema
// Created by chenguolin 2018-11-17
package user

// userTableModel user表结构定义
type userTableModel struct {
	ID         int64  `json:"id"`          //id
	UID        int64  `json:"uid"`         //uid
	Name       string `json:"name"`        //用户昵称
	Phone      string `json:"phone"`       //电话
	CreateTime int64  `json:"create_time"` //创建时间
	UpdateTime int64  `json:"update_time"` //更新时间
}
