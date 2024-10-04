package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nedssoft/go-basic-api/controllers"
)

type AuthRoutes struct {
	db     *gorm.DB
	router gin.RouterGroup
}

func NewAuthRoutes(router *gin.RouterGroup, db *gorm.DB) *AuthRoutes {
	return &AuthRoutes{
		db:     db,
		router: *router,
	}
}

func (r *AuthRoutes) RegisterRoutes() {
	authController := controllers.NewAuthController(r.db)
	r.router.POST("/login", authController.Login)
}
