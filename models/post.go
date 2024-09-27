package models

import (
	"time"
)

type Post struct {
	ID  uint `gorm:"primarykey"`
	Title string
	Body string `gorm:"type:text"`
	UserID uint 
	User User
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`
}
