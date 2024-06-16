package repo

import (
	"blog/dto"
	"blog/entities"
	"blog/model"
)

type (
	UserRepo interface {
		CreateUser(data model.User) error
		ReadUserByPhone(phone string) (model.UserInfo, error)
		ReadUserByID(id float64) (entities.User, error)
	}
	ArticleRepo interface {
		CreateArticle(data dto.ArticleDto, user entities.User) error
		ReadArticleByID(id int) (model.Article, error)
		ReadArticles(take int, skip int) ([]model.Article, error)
		ReadUserArticles(take int, skip int, userId int) ([]model.Article, error)
	}
)
