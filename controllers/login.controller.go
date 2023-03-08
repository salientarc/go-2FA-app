package controllers

import (
	"go-2fa-app/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ac *AuthController) LoginUser(ctx *gin.Context) {
	var payload *models.LoginUserInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Invalid email or Password",
			})
		return
	}

	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.Otp_enabled,
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
			"user":   userResponse,
		})
}
