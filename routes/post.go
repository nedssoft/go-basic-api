package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nedssoft/go-basic-api/controllers"
	"github.com/nedssoft/go-basic-api/middleware"
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
	r.router.GET("/posts", middleware.AuthMiddleware(r.db), postController.GetPosts)
	idRoute.GET("", middleware.AuthMiddleware(r.db), postController.GetPost)
	r.router.POST("/posts", middleware.AuthMiddleware(r.db), postController.CreatePost)
	idRoute.DELETE("", middleware.AuthMiddleware(r.db), postController.DeletePost)
	idRoute.PUT("", middleware.AuthMiddleware(r.db), postController.UpdatePost)
}