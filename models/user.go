package models

import (
	"time"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
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

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = hashPassword(user.Password)
	return
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}