package usecase

import (
	"blog/dto"
	"blog/entities"
	"blog/model"
	"blog/repo"
)

type articleUsecase struct {
	articleRepo repo.ArticleRepo
}

func NewArticleUsecase(articleRepo repo.ArticleRepo) Article {
	return &articleUsecase{
		articleRepo,
	}
}

func (au *articleUsecase) ReadArticleByID(id int) (model.Article, error) {
	return au.articleRepo.ReadArticleByID(id)
}

func (au *articleUsecase) CreateArticle(data dto.CreteArticleDto, user entities.User) error {
	article:=dto.ArticleDto{
		Title: data.Title,
		Content: data.Content,
	}
	return au.articleRepo.CreateArticle(article, user)
}

func (au *articleUsecase) ReadArticles(take int, skip int) ([]model.Article, error) {
	return au.articleRepo.ReadArticles(take, skip)
}

func (au *articleUsecase) ReadUserArticles(take int, skip int, userId int) ([]model.Article, error) {
	return au.articleRepo.ReadUserArticles(take, skip, userId)
}