package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-basic-api/auth"
	"github.com/nedssoft/go-basic-api/service"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		tokenGenerator := auth.NewJWTGenerator()
		userId, err := tokenGenerator.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		user, err := service.NewUserService(db).GetUserById(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()		
	}
}