package routes

import (
	"squadify-app/handlers"
	"squadify-app/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(api *gin.RouterGroup) {
	router := api.Group("/profile")

	router.Use(middleware.AuthMiddleware())
	{
		router.GET("", handlers.GetProfileByID)
		router.POST("", handlers.CreateProfile)
	}
}
