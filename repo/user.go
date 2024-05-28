package repo

import (
	"blog/model"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db,
	}
}

func (ur *userRepo) CreateUser(data model.User) error {
	err := ur.db.Create(data)
	if err != nil {
		return err.Error
	}
	return nil
}
