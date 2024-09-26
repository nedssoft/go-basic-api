package models

import (
	"time"
)

type User struct {
  ID        uint `gorm:"primarykey" json:"id"`
  CreatedAt *time.Time `json:"created_at,omitempty"`	
  UpdatedAt *time.Time `json:"updated_at,omitempty"`
  DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
	Posts []Post `json:"posts"`
}