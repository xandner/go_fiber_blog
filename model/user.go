package model

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Family   string `json:"family"`
	Password string `json:"password"`
}
