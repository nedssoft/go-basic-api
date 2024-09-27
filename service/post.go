package service

import (
	"github.com/nedssoft/go-basic-api/models"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.db.Create(post).Error
}

func (s *PostService) GetPost(id string) (*models.Post, error) {
	var post models.Post
	if err := s.db.Omit("user").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := s.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
