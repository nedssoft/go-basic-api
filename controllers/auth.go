package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/auth"
	"github.com/nedssoft/go-basic-api/data/requests"
	"github.com/nedssoft/go-basic-api/data/responses"
	"github.com/nedssoft/go-basic-api/service"
	"github.com/nedssoft/go-basic-api/utils"
	"gorm.io/gorm"
)

type AuthController struct {
	userService *service.UserService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{userService: service.NewUserService(db)}
}

func (c *AuthController) Login(gn *gin.Context) {
	var payload requests.LoginPayload
	if err := gn.BindJSON(&payload); err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errMessage := "Invalid email or password"
	user, err := c.userService.GetUserByEmail(payload.Email)
	if err != nil {
		gn.JSON(http.StatusUnauthorized, gin.H{"error": errMessage})
		return
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		gn.JSON(http.StatusUnauthorized, gin.H{"error": errMessage})
		return
	}
  tokenGenerator := auth.NewJWTGenerator()
	token, err := tokenGenerator.GenerateToken(user.ID)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
		return
	}

	gn.JSON(http.StatusOK, gin.H{"data": requests.LoginResponse{Token: token, User: responses.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}}})
}
