package repo

import (
	"blog/dto"
	"blog/entities"
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

func (a *articleRepo) CreateArticle(data dto.ArticleDto,user entities.User) error {
	newArticle:=model.Article{
		Title: data.Title,
		Content: data.Content,
		UserID: user.ID,
	}
	return a.db.Create(&newArticle).Error
}

func (a *articleRepo) ReadArticleByID(id int) (model.Article, error) {
	var article model.Article
	err := a.db.Where("id = ?", id).First(&article).Error
	return article, err
}