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

func (a *articleRepo) ReadArticles(take int, skip int) ([]model.Article, error) {
	var articles []model.Article
	err := a.db.Order("id DESC").Limit(take).Offset(skip).Find(&articles).Error
	return articles, err
}

func (a *articleRepo) ReadUserArticles(take int, skip int,userId int) ([]model.Article, error) {
	var articles []model.Article
	err := a.db.Where("user_id = ?",userId).Order("id DESC").Limit(take).Offset(skip).Find(&articles).Error
	return articles, err
}