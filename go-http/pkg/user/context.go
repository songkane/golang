/*
Package user 用户相关service context定义
Created by chenguolin 2018-11-17
*/
package user

import (
	init "gitlab.local.com/golang/httpserver/init"
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
		uRepo := newRepo(init.GetMysqlProxy())
		service = NewUserService(uRepo)
	})
}
