// Package user UserService
// Created by chenguolin 2018-05-03
package user

import (
	"time"

	"gitlab.local.com/golang/httpserver/bizerror"
)

// Service 结构体定义
type Service struct {
	uRepo *repo
}

// NewUserService 实例化一个UserService
func NewUserService(uRepo *repo) *Service {
	if uRepo == nil {
		panic("[Service - NewUserService] panic")
	}
	return &Service{
		uRepo: uRepo,
	}
}

// AddUserArgs AddUser函数参数
type AddUserArgs struct {
	UID   int64
	Name  string
	Phone string
}

// AddUser 新增一个用户
func (us *Service) AddUser(args *AddUserArgs) error {
	// 1. 参数校验
	if args == nil {
		return bizerror.ErrInvalidArguments
	}
	if args.UID <= 0 || args.Name == "" || args.Phone == "" {
		return bizerror.ErrInvalidArguments
	}

	// 2. userModel
	userModel := &userTableModel{
		UID:        args.UID,
		Name:       args.Name,
		Phone:      args.Phone,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	return us.uRepo.insertUser(userModel)
}

// DeleteUser 删除用户
func (us *Service) DeleteUser(uid int64) error {
	// TODO
	return us.uRepo.deleteUser(uid)
}

// UpdateUser 修改用户信息
func (us *Service) UpdateUser(uid int64, newName, newPhone string) error {
	if uid <= 0 || newName == "" || newPhone == "" {
		return bizerror.ErrInvalidArguments
	}

	return us.uRepo.updateUser(uid, newName, newPhone)
}

// InfoResult 查询结果定义
type InfoResult struct {
	UID        int64
	Name       string
	Phone      string
	CreateTime int64
	UpdateTime int64
}

// SelectUser 根据uid查询用户
func (us *Service) SelectUser(uid int64) (*InfoResult, error) {
	if uid <= 0 {
		return nil, bizerror.ErrInvalidArguments
	}

	userModel, err := us.uRepo.selectUser(uid)
	if err != nil {
		return nil, bizerror.ErrInvalidArguments
	}

	result := &InfoResult{
		UID:        userModel.UID,
		Name:       userModel.Name,
		Phone:      userModel.Phone,
		CreateTime: userModel.CreateTime,
		UpdateTime: userModel.UpdateTime,
	}

	return result, nil
}
