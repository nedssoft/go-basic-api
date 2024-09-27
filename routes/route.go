package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	db     *gorm.DB
	router gin.RouterGroup
}

func NewRoutes(router *gin.RouterGroup, db *gorm.DB) *Routes {
	return &Routes{
		db:     db,
		router: *router,
	}
}

func (r *Routes) RegisterRoutes() {
	postRoutes := NewPostRoutes(&r.router, r.db)
	postRoutes.RegisterRoutes()

	userRoutes := NewUserRoutes(&r.router, r.db)
	userRoutes.RegisterRoutes()
}
