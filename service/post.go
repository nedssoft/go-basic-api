package service

import (

	"github.com/nedssoft/go-basic-api/data/requests"
	"github.com/nedssoft/go-basic-api/data/responses"
	"github.com/nedssoft/go-basic-api/models"
	"gorm.io/gorm"
  "time"
  "github.com/nedssoft/go-basic-api/utils"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) CreatePost(payload *requests.PostPayload,userId uint) (postResponse *responses.PostResponse, err error) {
	post := models.Post{
		Title:  payload.Title,
		Body:   payload.Body,
		UserID: userId,
	}
	result := s.db.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}
  now := time.Now()
	return &responses.PostResponse{
		ID: post.ID,
		Title: post.Title,
		Body: post.Body,
		UserID: post.UserID,
		CreatedAt: utils.DefaultValue(post.CreatedAt, &now),
		UpdatedAt: utils.DefaultValue(post.UpdatedAt, &now),
	}, nil
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

func (s *PostService) DeletePost(id uint) error {
	return s.db.Delete(&models.Post{}, id).Error
}

func (s *PostService) UpdatePost(id string, payload *requests.PostUpdatePayload, userId uint) error {
	post := models.Post{
		Title:  payload.Title,
		Body:   payload.Body,
	}
	return s.db.Model(&models.Post{}).Where("id = ?", id).Where("user_id = ?", userId).Updates(&post).Error
}

