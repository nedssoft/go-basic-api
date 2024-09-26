package service

import (
	"github.com/nedssoft/learn-go/models"
	"gorm.io/gorm"
)

type UserService struct {
  db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
  return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
  return s.db.Create(user).Error
}

func (s *UserService) GetUser(id string) (*models.User, error) {
  var user models.User
  if err := s.db.Preload("Posts").First(&user, id).Error; err != nil {
    return nil, err
  }
  return &user, nil
}