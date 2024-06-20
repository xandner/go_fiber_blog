package usecase

import (
	"blog/dto"
	"blog/entities"
	"blog/model"
)

type (
	User interface {
		CreateUser(data model.User) error
		GetUserByID(id float64) (entities.User, error)
	}
	Auth interface {
		Login(loginData dto.UserLoginDto) (string,error)
		SignUp(userData model.User) error
	}
	Article interface {
		ReadArticleByID(id int) (model.Article, error)
		CreateArticle(data dto.CreteArticleDto,user entities.User) error
		ReadArticles(take int, skip int) ([]model.Article, error)
		ReadUserArticles(take int, skip int,userId int) ([]model.Article, error)
		UpdateArticle(articleId int, article model.Article) error
		DeleteArticle(articleId int) error
	}
)
