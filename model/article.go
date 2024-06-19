package model

type Article struct {
	Title  string `json:"title"`
	Content string `json:"content"`
	Published bool `json:"published"`
	UserID  uint   `json:"user_id"`
}