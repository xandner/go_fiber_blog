package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	Family string `gorm:"not null" json:"family"`
	Phone  string `gorn:"not null" json:"phone"`
}
