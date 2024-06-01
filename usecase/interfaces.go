package usecase

import (
	"blog/dto"
	"blog/model"
)

type (
	User interface {
		CreateUser(data model.User) error
	}
	Auth interface {
		Login(loginData dto.UserLoginDto) error
		SignUp(userData model.User) error
	}
)
