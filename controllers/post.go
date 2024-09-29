package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/service"
	"gorm.io/gorm"
	"github.com/nedssoft/go-basic-api/data/requests"
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
	var post requests.PostPayload
	if err := gn.BindJSON(&post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postResponse, err := c.PostService.CreatePost(&post)
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
	id := gn.Param("id")
	if err := c.PostService.DeletePost(id); err != nil {
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
	if err := c.PostService.UpdatePost(id, post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"post": post})
}
