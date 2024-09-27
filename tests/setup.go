package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	database "github.com/nedssoft/go-basic-api/bin/db"
)

func SetupTestRouter() (*gin.Engine, *gorm.DB) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Use in-memory SQLite for testing
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.Migrate(db)

	api := router.Group("/api/v1")
	routes := routes.NewRoutes(api, db)
	routes.RegisterRoutes()
	return router, db
}
