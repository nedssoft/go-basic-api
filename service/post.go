package service

import (
	"github.com/nedssoft/go-basic-api/data/requests"
	"github.com/nedssoft/go-basic-api/data/responses"
	"github.com/nedssoft/go-basic-api/models"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) CreatePost(payload *requests.PostPayload) error {
	post := models.Post{
		Title:  payload.Title,
		Body:   payload.Body,
		UserID: payload.UserID,
	}
	return s.db.Model(&models.Post{}).Create(&post).Error
}

func (s *PostService) GetPost(id string) (*responses.PostResponse, error) {
	var post responses.PostResponse
	if err := s.db.Model(&models.Post{}).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) GetPosts() ([]responses.PostResponse, error) {
	var posts []responses.PostResponse
	if err := s.db.Model(&models.Post{}).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) DeletePost(id string) error {
	return s.db.Delete(&models.Post{}, id).Error
}

func (s *PostService) UpdatePost(id string, payload *requests.PostPayload) error {
	post := models.Post{
		Title:  payload.Title,
		Body:   payload.Body,
		UserID: payload.UserID,
	}
	return s.db.Model(&models.Post{}).Where("id = ?", id).Updates(&post).Error
}
