// Package user service init
// Created by chenguolin 2018-11-17
package user

import (
	"gitlab.local.com/golang/httpserver/init"
)

var (
	service *Service
)

// GetUserService 获取UserService
func GetUserService() *Service {
	return service
}

func init() {
	init.AddInitFunc("UserService", func() {
		uRepo := newRepo(init.GetMysqlClient())
		service = NewUserService(uRepo)
	})
}
