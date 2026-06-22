package main

import (
	"go-image-api/config"
	_ "go-image-api/docs"
	"go-image-api/models"
	"go-image-api/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Image API
// @version 1.0
// @description Image Upload API with JWT + S3 + MySQL
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Image{},
	)

	r := gin.Default()

	routes.SetupRoutes(r)

	// Swagger
	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	r.Run(":8080")
}
