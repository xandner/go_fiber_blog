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

func (ur *userRepo) ReadUserByPhone(phone string) (model.User,error) {
	var user model.User
	err := ur.db.Where("phone = ?", phone).First(&user)
	if err != nil {
		return user,err.Error
	}
	return user,nil
}
