package dto

type UserDto struct {
	Name   string `json:"name" validate:"required"`
	Family string `json:"family" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}
