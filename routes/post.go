package routes

import (

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nedssoft/learn-go/controllers"
)

type PostRoutes struct {
	db *gorm.DB
	router gin.RouterGroup
}

func NewPostRoutes(router *gin.RouterGroup, db *gorm.DB) *PostRoutes {
	return &PostRoutes{
	  db:     db,
		router: *router,
	}
}

func (r *PostRoutes) RegisterRoutes() {
	postController := controllers.NewPostController(r.db)
	r.router.GET("/posts", postController.GetPosts)

	r.router.GET("/posts/:id", postController.GetPost)

	r.router.POST("/posts", postController.CreatePost)
}