package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/data/requests"
	"github.com/nedssoft/go-basic-api/service"
	"gorm.io/gorm"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		UserService: service.NewUserService(db),
	}
}

func (c *UserController) CreateUser(gn *gin.Context) {
	var user *requests.UserPayload
	if err := gn.ShouldBindJSON(&user); err != nil {

		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.HashPassword(); err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser, err := c.UserService.CreateUser(user)
	if  err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gn.JSON(http.StatusCreated, gin.H{"user": newUser})
}

func (c *UserController) GetUser(gn *gin.Context) {
	id := gn.Param("id")
	user, err := c.UserService.GetUser(id)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *UserController) GetUsers(gn *gin.Context) {
	users, err := c.UserService.GetUsers()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *UserController) DeleteUser(gn *gin.Context) {
	id := gn.Param("id")
	if err := c.UserService.DeleteUser(id); err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (c *UserController) UpdateUser(gn *gin.Context) {
	id := gn.Param("id")
	var user *requests.UserUpdatePayload
	if err := gn.ShouldBindJSON(&user); err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to extract user payload"})
		return
	}
	if err := c.UserService.UpdateUser(id, user); err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"user": user})
}
