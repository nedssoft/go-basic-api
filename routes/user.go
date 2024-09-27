package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/nedssoft/go-basic-api/controllers"
)

type UserRoutes struct {
	db *gorm.DB
	router gin.RouterGroup
}

func NewUserRoutes(router *gin.RouterGroup, db *gorm.DB) *UserRoutes {
	return &UserRoutes{
	  db:     db,
		router: *router,
	}
}

func (r *UserRoutes) RegisterRoutes() {
	userController := controllers.NewUserController(r.db)
	idRoute := r.router.Group("/users/:id")
	idRoute.GET("", userController.GetUser)
	r.router.GET("/users", userController.GetUsers)
	r.router.POST("/users", userController.CreateUser)
	idRoute.DELETE("", userController.DeleteUser)
	idRoute.PUT("", userController.UpdateUser)
}