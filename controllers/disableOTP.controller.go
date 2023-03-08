package controllers

import (
	"go-2fa-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac *AuthController) DisableOTP(ctx *gin.Context) {
	var payload *models.OTPInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		return
	}

	var user models.User

	result := ac.DB.First(&user, "id = ?", payload.UserId)
	if result.Error != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "User doesn't exist",
			})
		return
	}

	user.Otp_enabled = false
	ac.DB.Save(&user)

	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.Otp_enabled,
	}
	ctx.JSON(
		http.StatusOK, gin.H{
			"otp_disabled": true,
			"user":         userResponse,
		})
}
