package entities

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	Published bool `gorm:"default:false" json:"published"`  
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
}
