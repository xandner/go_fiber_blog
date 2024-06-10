package usecase

import (
	"blog/dto"
	"blog/model"
	"blog/repo"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func checkHashPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Printf("error: %v", err)
		return fmt.Errorf("password is incorrect")
	}
	return nil
}

func (au *authUsecase) Login(userData dto.UserLoginDto) (string,error) {
	user, err := au.userRepo.ReadUserByPhone(userData.Phone)
	if user.Phone != userData.Phone {
		return "",fmt.Errorf("user not found")
	}
	if err != nil {
		return "",err
	}
	err = checkHashPassword(userData.Password, user.Password)
	if err != nil {
		return "",err
	}
	// TODO create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.Id,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // 3 days
	})

	tokenString,err:=token.SignedString([]byte("secret"))
	if err != nil {
		return "",err
	}
	fmt.Println(tokenString)
	return tokenString,nil
}
func (au *authUsecase) SignUp(userData model.User) error {
	hashedPassword, err := hashPassword(userData.Password)
	if err != nil {
		return err
	}
	userData.Password = hashedPassword
	err = au.userRepo.CreateUser(userData)
	if err != nil {
		return err
	}
	return nil
}
