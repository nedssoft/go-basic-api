package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/models"
	"github.com/nedssoft/go-basic-api/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestRouter() (*gin.Engine, *gorm.DB) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Use in-memory SQLite for testing
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Post{}, &models.User{})

	api := router.Group("/api/v1")
	postRoutes := routes.NewPostRoutes(api, db)
	postRoutes.RegisterRoutes()
	userRoutes := routes.NewUserRoutes(api, db)
	userRoutes.RegisterUserRoutes()
	return router, db
}
