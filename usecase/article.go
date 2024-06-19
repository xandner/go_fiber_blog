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
	return au.articleRepo.CreateArticle(data, user)
}

func (au *articleUsecase) ReadArticles(take int, skip int) ([]model.Article, error) {
	return au.articleRepo.ReadArticles(take, skip)
}

func (au *articleUsecase) ReadUserArticles(take int, skip int, userId int) ([]model.Article, error) {
	return au.articleRepo.ReadUserArticles(take, skip, userId)
}

func (au *articleUsecase) UpdateArticle(articleId int, article model.Article) error {
	return au.articleRepo.UpdateArticle(articleId, article)
}