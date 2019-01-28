// Package user service instance
// Created by chenguolin 2018-11-17
package user

import (
	"gitlab.local.com/golang/go-http/instance"
)

var (
	service *Service
)

// GetUserService 获取UserService
func GetUserService() *Service {
	return service
}

func init() {
	instance.AddInitFunc("UserService", func() {
		uRepo := newRepo(instance.GetMysqlClient())
		service = NewUserService(uRepo)
	})
}
