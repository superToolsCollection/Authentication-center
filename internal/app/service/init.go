package service

import (
	"authentication-center/internal/app/model"
	"authentication-center/internal/app/service/user"
	"authentication-center/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)
//Init instantiate the service
func Init()  {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}