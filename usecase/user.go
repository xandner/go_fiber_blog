package usecase

import (
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
func (uu *userUsecase)CreateUser(user model.User) error {
	fmt.Println("USECASE")
	return uu.userRepo.CreateUser(user)
}
