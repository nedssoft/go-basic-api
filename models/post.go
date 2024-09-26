package models

import (
	"encoding/json"
	"time"
)

type Post struct {
	ID  uint `gorm:"primarykey" json:"id"`
	Title string `json:"title"`
	Body string `gorm:"type:text" json:"body"`
	UserID uint `json:"user_id"`
	User User `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (p Post) MarshalJSON() ([]byte, error) {
	type PostAlias Post
  return json.Marshal(PostAlias{
	  Title: p.Title,
		Body: p.Body,
		UserID: p.UserID,
		ID: p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
}