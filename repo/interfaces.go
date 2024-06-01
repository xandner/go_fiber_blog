package repo

import (
	"blog/model"
)

type (
	UserRepo interface {
		CreateUser(data model.User) error
		ReadUserByPhone()
	}
)
