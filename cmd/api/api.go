package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/routes"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	subrouter := router.Group("/api/v1")

	routes := routes.NewRoutes(subrouter, s.db)
	routes.RegisterRoutes()
	log.Println("Server running on port: ", s.addr)
	return router.Run()
}
