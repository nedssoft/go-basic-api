package models

import (
	"time"
)

type User struct {
  ID        uint `gorm:"primarykey"`
  CreatedAt *time.Time
  UpdatedAt *time.Time
  DeletedAt *time.Time `gorm:"index"`
	Password string
	Name string
	Email string
	Posts []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}