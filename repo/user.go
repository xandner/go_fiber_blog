package repo

import (
	"blog/entities"
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

func (ur *userRepo) ReadUserByPhone(phone string) (model.UserInfo, error) {
	var user entities.User
	err := ur.db.Where("phone = ?", phone).First(&user)
	foundedUser := model.UserInfo{
		Id:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Family:   user.Family,
		Password: user.Password,
	}
	if err != nil {
		return foundedUser, err.Error
	}
	return foundedUser, nil
}

func (ur *userRepo) ReadUserByID(id float64) (entities.User, error) {
	var user entities.User
	err := ur.db.Where("id = ?", id).First(&user)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}