package routes

import (
	"squadify-app/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(api *gin.RouterGroup) {
	router := api.Group("/auth")

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}
