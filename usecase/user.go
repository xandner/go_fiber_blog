package usecase

import (
	"blog/entities"
	"blog/model"
	"blog/repo"
	"fmt"
)

type userUsecase struct {
	userRepo repo.UserRepo
}

func NewUserUsecase(userRepo repo.UserRepo) User {
	return &userUsecase{
		userRepo,
	}
}
func (uu *userUsecase) CreateUser(data model.User) error {
	fmt.Println("USECASE")
	return uu.userRepo.CreateUser(data)
}

func (uu *userUsecase) GetUserByID(id float64) (entities.User, error) {
	return uu.userRepo.ReadUserByID(id)
}