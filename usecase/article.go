package usecase

import (
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