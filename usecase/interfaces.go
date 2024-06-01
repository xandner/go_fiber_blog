package usecase

import "blog/model"

type (
	User interface {
		CreateUser(data model.User) error
	}
	Auth interface {
		Login() error
		SignUp(userData model.User) error
	}
)
