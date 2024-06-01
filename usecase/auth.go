package usecase

import (
	"blog/model"
	"blog/repo"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo repo.UserRepo
}

func NewAuthUsecase(userRepo repo.UserRepo) Auth {
	return &authUsecase{
		userRepo,
	}
}

func (au *authUsecase) Login() error {
	fmt.Println("Login")
	return nil
}
func (au *authUsecase) hashPassword(password string) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func (au *authUsecase) SignUp(userData model.User) error {
	hashedPassword, err := au.hashPassword(userData.Password)
	if err != nil {
		return err
	}
	userData.Password = hashedPassword
	err=au.userRepo.CreateUser(userData)
	if err != nil {
		return err
	}
	return nil
}
