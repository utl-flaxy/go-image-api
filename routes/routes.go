package routes

import (
	"go-image-api/controllers"
	"go-image-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// 認証不要
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 認証必須
	auth := r.Group("/")

	auth.Use(middleware.AuthMiddleware())

	// 画像アップロード
	auth.POST("/images", controllers.UploadImage)

	// 画像一覧取得
	auth.GET("/images", controllers.GetImages)

	// 画像削除
	auth.DELETE("/images/:id", controllers.DeleteImage)
}
