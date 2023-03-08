package controllers

import (
	"go-2fa-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

func (ac *AuthController) ValidateOTP(ctx *gin.Context) {
	var payload *models.OTPInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		return
	}

	message := "Invalid token or user doesn't exist"

	var user models.User
	result := ac.DB.First(&user, "id = ?", payload.UserId)
	if result.Error != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": message,
			})
		return
	}

	valid := totp.Validate(payload.Token, user.Otp_secret)
	if !valid {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": message,
			})
		return
	}

	ctx.JSON(
		http.StatusOK, gin.H{
			"otp_valid": true,
		})
}
