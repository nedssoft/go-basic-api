package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/learn-go/models"
	"github.com/nedssoft/learn-go/service"
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
	var post *models.Post
	if err := gn.ShouldBindJSON(&post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}
  if err := c.PostService.CreatePost(post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post"})
		return
  }
	gn.JSON(http.StatusCreated, gin.H{"post": post})
}

func (c *PostController) GetPost(gn *gin.Context) {
	id := gn.Param("id")
  post, err := c.PostService.GetPost(id);
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get post"})
		return
  }
	gn.JSON(http.StatusOK, gin.H{"post": post})
}

func (c *PostController) GetPosts(gn *gin.Context)  {
	posts, err := c.PostService.GetPosts()
	if err != nil {
    log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get posts"})
		return
  }
	gn.JSON(http.StatusOK, gin.H{"posts": posts})
}