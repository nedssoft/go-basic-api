package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nedssoft/learn-go/models"
	"github.com/nedssoft/learn-go/service"
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

func (r *UserRoutes) RegisterUserRoutes() {
	userService := service.NewUserService(r.db)
	r.router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := userService.GetUser(id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})


	r.router.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to bind JSON",
			})
			return
		}
		log.Println(user)
		if err := userService.CreateUser(&user); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
		})
	})
}