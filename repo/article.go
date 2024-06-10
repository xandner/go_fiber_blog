package repo

import (
	"blog/model"

	"gorm.io/gorm"
)

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{
		db,
	}
}

func (a *articleRepo) CreateArticle(data model.Article) error {
	return a.db.Create(&data).Error
}

func (a *articleRepo) ReadArticleByID(id int) (model.Article, error) {
	var article model.Article
	err := a.db.First(&article, id).Error
	return article, err
}