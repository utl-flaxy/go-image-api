package controllers

import (
	"go-image-api/dto"
	"go-image-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary ユーザー登録
// @Description 新規ユーザーを登録する
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterRequest true "Register"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := services.Register(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "register failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

// Login godoc
// @Summary ログイン
// @Description JWTトークン取得
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Login"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	token, err := services.Login(
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "login failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
