package controllers

import (
	"log"
	"net/http"
	"path/filepath"

	"go-image-api/models"
	"go-image-api/repositories"
	"go-image-api/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadImage godoc
// @Summary 画像アップロード
// @Description S3へ画像をアップロードする
// @Tags Image
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param image formData file true "Image"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /images [post]
func UploadImage(c *gin.Context) {

	fileHeader, err := c.FormFile("image")
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "image required",
		})

		return
	}

	userID := c.MustGet("user_id").(uint)

	fileName :=
		uuid.New().String() +
			filepath.Ext(fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "open file failed",
		})

		return
	}

	defer file.Close()

	url, err := services.UploadToS3(
		file,
		fileName,
	)

	if err != nil {

		log.Println("===== S3 Upload Error =====")
		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	image := models.Image{
		UserID:   userID,
		FileName: fileName,
		FilePath: url,
	}

	err = repositories.CreateImage(image)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "db save failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "upload success",
		"file_name": fileName,
		"url":       url,
	})
}

// GetImages godoc
// @Summary 自分の画像一覧取得
// @Description ログインユーザーの画像一覧取得
// @Tags Image
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /images [get]
func GetImages(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	images, err :=
		repositories.GetImagesByUserID(userID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed",
		})

		return
	}

	c.JSON(http.StatusOK, images)
}

// DeleteImage godoc
// @Summary 画像削除
// @Description S3とDBから画像を削除する
// @Tags Image
// @Produce json
// @Security BearerAuth
// @Param id path int true "Image ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /images/{id} [delete]
func DeleteImage(c *gin.Context) {

	id := c.Param("id")

	image, err := repositories.FindImageByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "image not found",
		})

		return
	}

	err = services.DeleteFromS3(
		image.FilePath,
	)

	if err != nil {

		log.Println("===== S3 Delete Error =====")
		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = repositories.DeleteImage(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "db delete failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
