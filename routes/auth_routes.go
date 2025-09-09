package routes

import (
	"net/http"
	"squadify-app/handlers"
	"squadify-app/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(api *gin.RouterGroup) {
	router := api.Group("/auth")

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.Use(middleware.AuthMiddleware())
	{
		router.GET("/profile", func(c *gin.Context) {
			email, _ := c.Get("email")
			c.JSON(http.StatusOK, gin.H{"message": "Welcome!", "email": email})
		})
	}
}
