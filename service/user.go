package service

import (
	"github.com/nedssoft/go-basic-api/data/requests"
	"github.com/nedssoft/go-basic-api/data/responses"
	"github.com/nedssoft/go-basic-api/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(payload *requests.UserPayload) (userResponse *responses.UserResponse, err error) {
	user := models.User{
		Name: payload.Name,
		Email: payload.Email,
	}
	result := s.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &responses.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) GetUser(id string) (*responses.UserPostsResponse, error) {
	var user models.User
	if err := s.db.Model(&models.User{}).Preload("Posts").First(&user, id).Error; err != nil {
		return nil, err
	}
  userPostsResponse := responses.UserPostsResponse{
    UserResponse:responses.UserResponse{
      ID: user.ID,
      Name: user.Name,
      Email: user.Email,
      CreatedAt: user.CreatedAt,
      UpdatedAt: user.UpdatedAt,
    },
    Posts: make([]responses.PostResponse, len(user.Posts)),
  }
  for i, post := range user.Posts {
    userPostsResponse.Posts[i] = responses.PostResponse{
      ID: post.ID,
      Title: post.Title,
      Body: post.Body,
      UserID: post.UserID,
      CreatedAt: post.CreatedAt,
      UpdatedAt: post.UpdatedAt,
    }
  }
	return &userPostsResponse, nil
}

func (s *UserService) GetUsers() ([]responses.UserResponse, error) {
	var users []responses.UserResponse
	if err := s.db.Model(&models.User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) DeleteUser(id string) error {
	return s.db.Delete(&models.User{}, id).Error
}

func (s *UserService) UpdateUser(id string, user *requests.UserUpdatePayload) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}
