package repo

import (
	"blog/model"
)

type (
	UserRepo interface {
		CreateUser(data model.User) error
		ReadUserByPhone(phone string) (model.UserInfo, error)
	}
	ArticleRepo interface {
		CreateArticle(data model.Article) error
		ReadArticleByID(id int) (model.Article, error)
	}
)
