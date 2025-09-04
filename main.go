package main

import (
	"fmt"
	"os"
	"squadify-app/config"
	"squadify-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	PORT := os.Getenv("PORT")

	config.ConnectDB()
	// utils.InitJWT()

	app := gin.Default()
	routes.RegisterRoutes(app)

	app.Run(fmt.Sprintf(":%v", PORT))
}
