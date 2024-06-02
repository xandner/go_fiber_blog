package model

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Family   string `json:"family"`
	Password string `json:"password"`
}


type UserInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Family string `json:"family"`
	Password string `json:"password"`
}