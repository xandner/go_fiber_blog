package model

type Article struct {
	Title  string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}