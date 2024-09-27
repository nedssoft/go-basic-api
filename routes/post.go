package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nedssoft/go-basic-api/controllers"
)

type PostRoutes struct {
	db     *gorm.DB
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
	idRoute := r.router.Group("/posts/:id")
	r.router.GET("/posts", postController.GetPosts)
	idRoute.GET("", postController.GetPost)
	r.router.POST("/posts", postController.CreatePost)
	idRoute.DELETE("", postController.DeletePost)
	idRoute.PUT("", postController.UpdatePost)
}