package usecase

import (
	"blog/repo"
	"fmt"
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

func (au *authUsecase) SignUp() error {
	fmt.Println("SignUp")
	return nil
}
