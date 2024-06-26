package dto

type ArticleDto struct {
	Title  string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type CreteArticleDto struct {
	Title  string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Published bool `json:"published" validate:"required"`
}