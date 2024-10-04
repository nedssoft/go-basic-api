package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/data/requests"
	"github.com/nedssoft/go-basic-api/models"
	"github.com/nedssoft/go-basic-api/service"
	"gorm.io/gorm"
)

type PostController struct {
	PostService *service.PostService
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		PostService: service.NewPostService(db),
	}
}

func (c *PostController) CreatePost(gn *gin.Context) {
	user, ok := gn.MustGet("user").(*models.User)
	if !ok {
		log.Println("user not found")
		gn.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	var post requests.PostPayload
	if err := gn.BindJSON(&post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postResponse, err := c.PostService.CreatePost(&post, user.ID)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post"})
		return
	}
	gn.JSON(http.StatusCreated, gin.H{"post": postResponse})
}

func (c *PostController) GetPost(gn *gin.Context) {
	id := gn.Param("id")
	post, err := c.PostService.GetPost(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get post"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"post": post})
}

func (c *PostController) GetPosts(gn *gin.Context) {
	posts, err := c.PostService.GetPosts()
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get posts"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (c *PostController) DeletePost(gn *gin.Context) {
	user, ok := gn.MustGet("user").(*models.User)
	if !ok {
		log.Println("user not found")
		gn.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	id := gn.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	post, err := c.PostService.GetPost(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get post"})
		return
	}
	if post.UserID != user.ID {
		gn.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this post"})
		return
	}
	if err := c.PostService.DeletePost(uint(uid)); err != nil {
		log.Println(err)
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func (c *PostController) UpdatePost(gn *gin.Context) {
	id := gn.Param("id")
	var post *requests.PostUpdatePayload
	if err := gn.BindJSON(&post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, ok := gn.MustGet("user").(*models.User)
	if !ok {
		log.Println("user not found")
		gn.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	if err := c.PostService.UpdatePost(id, post, user.ID); err != nil {
		log.Println(err)
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"post": post})
}
